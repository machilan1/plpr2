## confirm: Confirm to execute
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
warning-main-origin:
	@echo "=================================================================================";\
	echo "⚠️  WARNING: F-O-R-C-E main branch on local machine aligned with latest origin...";

## explain-basic-rebase: Explain the whole basic-rebase process.
explain-basic-rebase:
	@echo "================================================================================================================================";\
	echo "The basic-rebase workflow assumes that you're already on the feat branch, had made some code changes and want to merge with main";\
	echo "branch using rebase.";\
	echo "";\
	echo "Step1. Confirm you are NOT on the main branch.";\
	echo "Step2. Add and commit changes on LOCAL feat branch at LOCAL machine. This can be done multiple times before next step.";\
	echo "Step3. F-O-R-C-E LOCAL main branch align with REMOTE. ONLY for working directory cleaning, we still use origin/main to rebase.";\
	echo "Step4. Confirm feat branch status between REMOTE and LOCAL, we may need to rebase or resolve conflicts.";\
	echo "Step5. Rebase (force-with-lease) current branch with REMOTE main branch.";\
	echo "=================================================================================================================================";

## explain-basic-branch: Explain the whole basic-branch process.
explain-basic-branch:
	@echo "================================================================================================================================";\
    	echo "The basic-branch workflow assumes that you've just git clone the project and current branch is main.";\
    	echo "";\
    	echo "Step1. Confirm you are on the main branch.";\
    	echo "Step2. Add and new branch and make it upstream with origin.";\
    	echo "=================================================================================================================================";

## basic-branch-1: Confirm LOCAL main and REMOTE main sync.
basic-branch-1:
	@git fetch origin
		@if git diff --quiet main origin/main; then \
			echo "===================================================================";\
			echo "✅  LOCAL Main Branch is sync with REMOTE origin.";\
			echo "===================================================================";\
		else \
		  	echo "===================================================================";\
			echo "❌  LOCAL Main Branch is NOT sync with REMOTE origin. Please check.";\
			echo "===================================================================";\
		fi

## basic-branch-2: Create a new branch for editing.
basic-branch-2:
	@read -p "Please enter the branch name you want to add: " Branchname;\
	if [ -z "$$Branchname" ]; then \
		echo "==================================================================";\
		echo "❌ No branch name provided, aborting.";\
		echo "==================================================================";\
	else \
		echo "git checkout -b \"$$Branchname\"...";\
		git checkout -b "$$Branchname";\
		echo "set \"$$Branchname\" upstream with origin...";\
		git push --set-upstream origin "$$Branchname";\
		echo "==================================================================";\
		echo "✅  A brand new branch is created, you can start editing now.";\
		echo "==================================================================";\
	fi

## basic-rebase-1: Confirm current LOCAL branch is not main.
basic-rebase-1: confirm
	@echo 'git check current branch...'
	@current_branch=$$(git rev-parse --abbrev-ref HEAD);\
	if [ "$$current_branch" = "main" ] || [ "$$current_branch" = "master" ]; then \
		echo "================================================================================================";\
		echo "❌  You are on the main or master branch, please checkout to your new branch before any actions.";\
		echo "================================================================================================";\
		exit 1;\
	else \
	  	echo "=========================================================================================================";\
		echo "✅  Confirmed you are NOT on main or master branch, please confirm this is the branch you want to rebase.";\
		echo "=========================================================================================================";\
	fi

## basic-rebase-2: Commit LOCAL feat branch changes. This can be done multiple times before next step.
basic-rebase-2: basic-rebase-1
	@read -p "Please enter the commit message of local branch change: " localCommitMessage;\
	if [ -z "$$localCommitMessage" ]; then \
		echo "==================================================================";\
		echo "❌ No local commit message provided, aborting.";\
		echo "==================================================================";\
	else \
		echo "git commit -m \"$$localCommitMessage\"...";\
		git add .;\
		git commit -m "$$localCommitMessage";\
		echo "====================================================================";\
		echo "✅  Changes are commited in the LOCAL branch, or already the latest.";\
		echo "====================================================================";\
	fi

## basic-rebase-3: WARNING,F-O-R-C-E main on LOCAL machine aligned with REMOTE origin. test
basic-rebase-3: check-commit warning-main-origin confirm
	@echo '⚠️  Warning: Force main branch on local machine to align with the latest origin...'
	@current_branch=$$(git rev-parse --abbrev-ref HEAD);\
	read -p "Please enter the name of the main branch (empty if 'main'): " mainName;\
	if [ -z "$$mainName" ]; then \
		mainName="main";\
	fi;\
	if [ "$$mainName" != "main" ] && [ "$$mainName" != "master" ]; then \
      echo "Error: The main branch name must be 'main' or 'master'. Aborting.";\
      exit 1;\
    fi;\
	echo "Switching to \"$$mainName\"...";\
	git checkout "$$mainName";\
	echo "Fetching origin...";\
	git fetch origin;\
	if git diff --quiet $$mainName origin/$$mainName; then \
	  	echo "==================================================================";\
		echo "✅  Local $$mainName is already up to date with origin/$$mainName.";\
		echo "==================================================================";\
	else \
		echo "Resetting $$mainName to align with origin/$$mainName...";\
		git reset --hard origin/$$mainName;\
		echo "============================================================================";\
		echo "✅  Force reset $$mainName completed. $$mainName is now aligned with origin.";\
		echo "============================================================================";\
	fi;\
	echo "Switching back to the original branch...";\
	git checkout "$$current_branch";

## basic-rebase-4: Confirm feat branch head on LOCAL machine is aligned REMOTE origin, if not please resolve it before further actions.
basic-rebase-4: check-commit basic-rebase-1
	@BRANCH_NAME=$$(git rev-parse --abbrev-ref HEAD);\
		REMOTE_BRANCH=origin/$$BRANCH_NAME;\
		echo "Checking the relationship between local and remote branches: $$BRANCH_NAME and $$REMOTE_BRANCH";\
		result=`git rev-list --left-right --count $$REMOTE_BRANCH...$$BRANCH_NAME`;\
		echo "Remote ahead: $$(echo $$result | cut -d' ' -f1) commits";\
		echo "Local ahead: $$(echo $$result | cut -d' ' -f2) commits";\
		if [ "$$(echo $$result | cut -d' ' -f1)" -eq "0" ] && [ "$$(echo $$result | cut -d' ' -f2)" -eq "0" ]; then \
			echo "==================================================================";\
			echo "✅  LOCAL/REMOTE Branches are synchronized, no differences.";\
			echo "==================================================================";\
		elif [ "$$(echo $$result | cut -d' ' -f1)" -gt "0" ] && [ "$$(echo $$result | cut -d' ' -f2)" -eq "0" ]; then \
			echo "=====================================================================================";\
			echo "⚠️  REMOTE branch is ahead, consider running 'git pull origin $$BRANCH_NAME' to sync.";\
			echo "=====================================================================================";\
		elif [ "$$(echo $$result | cut -d' ' -f2)" -gt "0" ] && [ "$$(echo $$result | cut -d' ' -f1)" -eq "0" ]; then \
			echo "==================================================================";\
			echo "💡  LOCAL branch is ahead, run next step";\
			echo "==================================================================";\
		else \
			echo "=================================================================================";\
			echo "⚠️⚠️⚠️  Branches have diverged, you may need to rebase or resolve conflicts.";\
			echo "=================================================================================";\
		fi

## basic-rebase-5: Rebase current branch with REMOTE feat branch (force-with-lease).
basic-rebase-5: check-commit
		@current_branch=$$(git rev-parse --abbrev-ref HEAD);\
		read -p "Please enter the name of the main branch (empty if 'main'): " mainName;\
		if [ -z "$$mainName" ]; then \
			mainName="main";\
		fi;\
		if [ "$$mainName" != "main" ] && [ "$$mainName" != "master" ]; then \
          echo "Error: The main branch name must be 'main' or 'master'. Aborting.";\
          exit 1;\
        fi;\
		echo "Rebase \"$$mainName\" to main/master...";\
		git rebase origin/$$mainName;\
		echo "==================================================================";\
		echo "push (force with lease) rebased branch \"$$mainName\" to origin...";\
		echo "==================================================================";\
		git push --force-with-lease origin $$current_branch;\
		echo "==================================================================";\
		echo "✅  All steps are done, you can now create a Pull Request.";\
		echo "==================================================================";

## explain-basic-rebase-detailed: Visualization.
explain-basic-rebase-detailed:
	@echo "❌  Something should not happen.";\
	echo "✅  Something happened here.";\
	echo "⚠️ Something should be noticed.";\
	echo "===============================================================================================================================";\
	echo "Step1. Confirm you are not in the main branch.";\
	echo "";\
	echo "LOCAL:";\
	echo "Main1 ----> Main2 (❌  You should be in feat branch)";\
	echo "                 \\";\
	echo "                  ----> feat-branch (✅  You should be in this branch)";\
	echo "";\
	echo "REMOTE(origin):";\
	echo "Main1 ----> Main2";\
	echo "                 \\";\
	echo "                  ----> feat-branch (⚠️ Upstream should be set already, could have 0~N commits.)";\
	echo "===============================================================================================================================";\
	echo "Step2. Add and commit changes on LOCAL feat branch at LOCAL machine. This can be done multiple times before next step.";\
	echo "";\
	echo "LOCAL:";\
	echo "Main1 ----> Main2";\
	echo "                 \\";\
	echo "                  ----> feat-branch-1 (✅  commit changes here, you can commit multiple times.)";\
	echo "";\
	echo "REMOTE(origin):";\
	echo "Main1 ----> Main2 ----> Main3 (⚠️ Someone may change main through PR review when you change local codes.)";\
	echo "                 \\";\
	echo "                  ----> feat-branch";\
	echo "===============================================================================================================================";\
	echo "Step3. F-O-R-C-E LOCAL main branch align with REMOTE. ONLY for working directory cleaning, we still use origin/main to rebase.";\
	echo "";\
	echo "LOCAL:";\
	echo "Main1 ----> Main2 ----> Main3 (✅  FORCE local main branch align with REMOTE main)";\
	echo "                 \\";\
	echo "                  ----> feat-branch-1";\
	echo "";\
	echo "REMOTE(origin):";\
	echo "Main1 ----> Main2 ----> Main3 (⚠️ Someone may change main through PR review when you change local codes.)";\
	echo "                 \\";\
	echo "                  ----> feat-branch";\
	echo "===============================================================================================================================";\
	echo "Step4. Confirm feat branch status between REMOTE and LOCAL, we may need to rebase or resolve conflicts.";\
	echo "";\
	echo "LOCAL:";\
	echo "Main1 ----> Main2 ----> Main3";\
	echo "                 \\";\
	echo "                  ----> feat-branch-1(Me, step2) ----> feat-branch-2(Me, step2)";\
	echo "";\
	echo "REMOTE(origin):";\
	echo "Main1 ----> Main2 ----> Main3";\
	echo "                 \\";\
	echo "                  ----> feat-branch-1(Me, pushed) ----> feat-branch-X(⚠️ Someone, pushed)";\
	echo "";\
	echo "Above origin shows conflict condition that you need to resolve. If LOCAL branch simply ahead REMOTE, you can go rebase(next step).";\
	echo "(We will not show conflict in next step)";\
	echo "===============================================================================================================================";\
	echo "Step5. Rebase (force-with-lease) current LOCAL branch with REMOTE main branch, and push to REMOTE. If someone pushs update as ";\
	echo "above conflict, force-with-lease will still fails, waiting resolving the conflict.";\
	echo "";\
	echo "LOCAL:";\
	echo "Main1 ----> Main2 ----> Main3";\
	echo "                 ";\
	echo "                         LOCAL: ----> feat-branch-1(Me, step2) ----> feat-branch-2(Me, step2)";\
	echo "                               /                                                            \\";\
	echo "                              / ✅  Rebase LOCAL branch                                       \\  ✅  Push rebased LOCAL feat branch to REMOTE";\
	echo "REMOTE(origin):              /  (with REMOTE origin, prevent someone update AGAIN.)           \\";\
	echo "Main1 ----> Main2 ----> Main3                                                                  \\";\
	echo "                                                                                         REMOTE(origin) feat branch ---->---";\
	echo "                                                                                                  ";\
	echo "";\

br-back: check-commit
	@git checkout main;\
	git fetch origin;\
	git pull;\

br-1: basic-rebase-1
br-2: basic-rebase-2
br-3: basic-rebase-3
br-4: basic-rebase-4
br-5: basic-rebase-5
bb-1: basic-branch-1
bb-2: basic-branch-2
br-b: br-back

check-commit:
	@echo "Checking if there are uncommitted changes...";\
	if [ -n "$$(git status --porcelain)" ]; then \
		echo "There are uncommitted changes!"; \
		exit 1;\
	else \
		echo "✅ Working directory is clean."; \
	fi