# Go Tool

A generic Go project template, meant to be used as a starting point for new projects.

## Pre-requisites

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

## Start your new project

1. Create a new project using `gonew`:

```sh
GOPRIVATE=github.com/mayainfo gonew github.com/mayainfo/gotool github.com/mydomain/myproject
```

2. Initialize the project:

```sh
cd myproject && make tidy && git init
```

Happy coding!

## Backend guide

### Install dependencies / packages

Install project dependencies & pre-commit hooks.

```sh
make dev-brew
```

Install Go tools, including formatting, linting, and testing tools.

```sh
make dev-gotooling
```

### Setup the environment

```sh
cp .env.example .env
```

Now you can start all the services by running:

```sh
make dev-up
```

To stop and remove the containers, run:

```sh
make dev-down
```

### Development with hot reload

```sh
air
```

### Run the tests

```sh
make test
```

## Frontend support

For frontend support, we will use Nx as the monorepo tool. The following steps will guide you through the process of setting up a new Nx workspace.

1. Create a new Nx workspace

In the root of the repository, run the following command:

```sh
export WORKSPACE_NAME=org &&
npx create-nx-workspace@latest --name=${WORKSPACE_NAME} --appName=app --packageManager=pnpm --nxCloud=skip --skipGit=true --preset=angular-monorepo
```

2. Clean up unnecessary files

Remove the app that was created by default

```sh
export WORKSPACE_NAME=org &&
cd ${WORKSPACE_NAME} &&
nx g @nx/workspace:remove --projectName=app &&
cd ..
```

Remove the .gitignore file, lock file, .nx folder, .angular folder, and the node_modules folder

```sh
export WORKSPACE_NAME=org &&
rm -rf ${WORKSPACE_NAME}/README.md \
       ${WORKSPACE_NAME}/.gitignore \
       ${WORKSPACE_NAME}/pnpm-lock.yaml \
       ${WORKSPACE_NAME}/.nx \
       ${WORKSPACE_NAME}/.angular \
       ${WORKSPACE_NAME}/node_modules
```

3. Move the files inside <WORKSPACE_NAME> to the root of the repository

```sh
export WORKSPACE_NAME=org &&
mv ${WORKSPACE_NAME}/* . &&
rm -rf ${WORKSPACE_NAME}
```

4. Customize your workspace

- Update the `name` field in the `package.json` file. This will be used as prefix for further projects created in the workspace.

```json
{
  "name": "@<WORKSPACE_NAME>/source"
}
```

5. Install the necessary dependencies

```sh
pnpm i
```

6. Create a new project

```sh
nx g @nx/angular:app web/apps/<PROJECT_NAME> --style=css --bundler=esbuild
```

7. Create a new library

```sh
nx g @nx/angular:lib web/libs/<LIBRARY_NAME>
```
