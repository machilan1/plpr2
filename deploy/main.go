package main

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/cloudbuild"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/cloudrun"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/cloudrunv2"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/secretmanager"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/sql"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// TODO: Cloud scheduler for worker
// TODO: Set deletion protection to true for production resources (maybe use env config for pulumi)

const DB_STORAGE_ABBR = "plpc2"
const REPO_NAME = "plpr2"
const REPO_OWNER = "machilan1"
const DEFAULT_DOMAIN = "example.com"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// -------------------------------------------------------------------------------------------------------------
		// Config

		env := config.Require(ctx, "env")
		adminEmail := config.Require(ctx, "admin-email")

		gcpCfg := config.New(ctx, "gcp")
		gcpProject := gcpCfg.Require("project")

		// -------------------------------------------------------------------------------------------------------------
		// API services

		servicesToEnable := []string{
			"artifactregistry.googleapis.com",
			"cloudbuild.googleapis.com",
			"secretmanager.googleapis.com",
			"sqladmin.googleapis.com",
			"run.googleapis.com",
			"cloudscheduler.googleapis.com",
		}
		enabledService := make([]pulumi.Resource, len(servicesToEnable))
		for i, service := range servicesToEnable {
			srv, err := projects.NewService(ctx, fmt.Sprintf("enable-%s", service), &projects.ServiceArgs{
				Service:          pulumi.String(service),
				DisableOnDestroy: pulumi.Bool(false),
			})
			if err != nil {
				return err
			}
			enabledService[i] = srv
		}

		// -------------------------------------------------------------------------------------------------------------
		// Service account

		devopsSA, err := serviceaccount.NewAccount(ctx, "devops-sa", &serviceaccount.AccountArgs{
			AccountId:   pulumi.String("devops"),
			DisplayName: pulumi.String("Service account for DevOps"),
		})
		if err != nil {
			return err
		}
		devopsSARoles := []string{
			"roles/secretmanager.admin",
			"roles/logging.logWriter",       // for writing logs to Cloud Logging (e.g. Cloud Build logs)
			"roles/artifactregistry.writer", // for pushing images to Artifact Registry (e.g. during Cloud Build)
			"roles/run.admin",               // for deploying Cloud Run services
			"roles/iam.serviceAccountUser",  // for attaching a service account to a resource (e.g. Cloud Run)
			"roles/cloudsql.client",         // for connecting to Cloud SQL
			"roles/cloudsql.instanceUser",   // for login as IAM user to Cloud SQL
			"roles/cloudscheduler.admin",    // for scheduling jobs with Cloud Scheduler
		}
		for _, role := range devopsSARoles {
			_, err = projects.NewIAMMember(ctx, fmt.Sprintf("devops-sa-iam-%s", role), &projects.IAMMemberArgs{
				Project: pulumi.String(gcpProject),
				Role:    pulumi.String(role),
				Member:  devopsSA.Member,
			})
			if err != nil {
				return err
			}
		}

		webSrvcSA, err := serviceaccount.NewAccount(ctx, "web-service-sa", &serviceaccount.AccountArgs{
			AccountId:   pulumi.String("web-service"),
			DisplayName: pulumi.String("Service account for Cloud Run web service"),
		})
		if err != nil {
			return err
		}
		// Intentionally don't grant any roles to the web service account due to the least privilege principle.

		apiSrvcSA, err := serviceaccount.NewAccount(ctx, "api-service-sa", &serviceaccount.AccountArgs{
			AccountId:   pulumi.String("api-service"),
			DisplayName: pulumi.String("Service account for Cloud Run api service"),
		})
		if err != nil {
			return err
		}
		apiAcctRoles := []string{
			"roles/storage.objectAdmin",            // for managing objects in Cloud Storage
			"roles/iam.serviceAccountTokenCreator", // for signing URL
			"roles/cloudsql.client",                // for connecting to Cloud SQL
			"roles/cloudsql.instanceUser",          // for login as IAM user to Cloud SQL
			"roles/secretmanager.secretAccessor",
		}
		for _, role := range apiAcctRoles {
			_, err = projects.NewIAMMember(ctx, fmt.Sprintf("api-service-sa-iam-%s", role), &projects.IAMMemberArgs{
				Project: pulumi.String(gcpProject),
				Role:    pulumi.String(role),
				Member:  apiSrvcSA.Member,
			})
			if err != nil {
				return err
			}
		}

		adminAcctRoles := []string{
			"roles/writer",
		}
		for _, role := range adminAcctRoles {
			_, err = projects.NewIAMMember(ctx, fmt.Sprintf("admin-sa-iam-%s", role), &projects.IAMMemberArgs{
				Project: pulumi.String(gcpProject),
				Role:    pulumi.String(role),
				Member:  pulumi.String("user:" + adminEmail),
			})
			if err != nil {
				return err
			}
		}

		// -------------------------------------------------------------------------------------------------------------
		// Secret Manager

		sessionSecretKey, err := random.NewRandomBytes(ctx, "secret-session-data", &random.RandomBytesArgs{
			Length: pulumi.Int(64),
		})
		if err != nil {
			return err
		}
		sessionSecret, err := secretmanager.NewSecret(ctx, "secret-session", &secretmanager.SecretArgs{
			SecretId: pulumi.String("session-key"),
			Replication: &secretmanager.SecretReplicationArgs{
				Auto: &secretmanager.SecretReplicationAutoArgs{},
			},
		},
			pulumi.DependsOn(enabledService),
		)
		if err != nil {
			return err
		}
		_, err = secretmanager.NewSecretVersion(ctx, "secret-session-version", &secretmanager.SecretVersionArgs{
			Secret:     sessionSecret.ID(),
			SecretData: sessionSecretKey.Base64,
		})
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Artifact registry

		ar, err := artifactregistry.NewRepository(ctx, "artifact-registry", &artifactregistry.RepositoryArgs{
			RepositoryId: pulumi.String("cloud-run-source-deploy"),
			Location:     pulumi.String("asia-east1"),
			Description:  pulumi.String("docker repository for cloud run"),
			Format:       pulumi.String("DOCKER"),
			DockerConfig: &artifactregistry.RepositoryDockerConfigArgs{
				ImmutableTags: pulumi.Bool(false), // TODO: might need to change this
			},
			//TODO: investigate these
			//CleanupPolicyDryRun: pulumi.Bool(false),
			//CleanupPolicies: &artifactregistry.RepositoryCleanupPolicyArray{
			//},
		},
			pulumi.DependsOn(enabledService),
		)
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Cloud Storage

		bucket, err := storage.NewBucket(ctx, "main-bucket", &storage.BucketArgs{
			Name:                     pulumi.String(DB_STORAGE_ABBR + "-" + env),
			Location:                 pulumi.String("ASIA-EAST1"),
			ForceDestroy:             pulumi.Bool(!isProdEnv(env)),
			UniformBucketLevelAccess: pulumi.Bool(true),
			PublicAccessPrevention:   pulumi.String("enforced"),
			StorageClass:             pulumi.String("STANDARD"),
			Cors: storage.BucketCorArray{
				&storage.BucketCorArgs{
					Origins: pulumi.StringArray{
						pulumi.String("*"),
					},
					Methods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("PUT"),
						pulumi.String("OPTIONS"),
					},
					ResponseHeaders: pulumi.StringArray{
						pulumi.String("Content-Type"),
						// pulumi.String("x-goog-acl"), // enable this if you want to set ACL (TODO: investigate if this has effect when using uniform bucket-level access)
					},
					MaxAgeSeconds: pulumi.Int(3600),
				},
			},
		})
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Cloud SQL

		instance, err := sql.NewDatabaseInstance(ctx, "db-instance", &sql.DatabaseInstanceArgs{
			Name:            pulumi.String(DB_STORAGE_ABBR),
			Region:          pulumi.String("asia-east1"),
			DatabaseVersion: pulumi.String("POSTGRES_17"),
			Settings: &sql.DatabaseInstanceSettingsArgs{
				Edition: pulumi.String("ENTERPRISE"),

				// Tier: pulumi.String("db-f1-micro"), // 0.2 vCPU, 0.6 GB RAM
				// Tier: pulumi.String("db-g1-small"), // 0.5 vCPU, 1.7 GB RAM
				Tier: pulumi.String("db-custom-1-3840"), // 1 vCPU, 3.75 GB RAM

				IpConfiguration: &sql.DatabaseInstanceSettingsIpConfigurationArgs{
					SslMode: pulumi.String("ENCRYPTED_ONLY"),
				},
				BackupConfiguration: &sql.DatabaseInstanceSettingsBackupConfigurationArgs{
					BackupRetentionSettings: &sql.DatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsArgs{
						RetainedBackups: pulumi.Int(7),
						RetentionUnit:   pulumi.String("COUNT"),
					},
					Enabled:                     pulumi.Bool(true),
					PointInTimeRecoveryEnabled:  pulumi.Bool(true),
					TransactionLogRetentionDays: pulumi.Int(7),
					StartTime:                   pulumi.String("19:00"), // UTC time. Equivalent to 3:00-7:00 AM in Asia/Taipei.
				},
				DiskAutoresize: pulumi.Bool(true),
				InsightsConfig: &sql.DatabaseInstanceSettingsInsightsConfigArgs{
					QueryInsightsEnabled: pulumi.Bool(true),
				},
				AvailabilityType: pulumi.String("ZONAL"),
				DatabaseFlags: sql.DatabaseInstanceSettingsDatabaseFlagArray{
					&sql.DatabaseInstanceSettingsDatabaseFlagArgs{
						Name:  pulumi.String("cloudsql.iam_authentication"),
						Value: pulumi.String("on"),
					},
					&sql.DatabaseInstanceSettingsDatabaseFlagArgs{
						Name:  pulumi.String("cloudsql.enable_pgaudit"),
						Value: pulumi.String("on"),
					},
				},
				PasswordValidationPolicy: &sql.DatabaseInstanceSettingsPasswordValidationPolicyArgs{
					DisallowUsernameSubstring: pulumi.Bool(true),
					EnablePasswordPolicy:      pulumi.Bool(true),
					ReuseInterval:             pulumi.Int(2),
					Complexity:                pulumi.String("COMPLEXITY_DEFAULT"), // Must contain an uppercase letter, a lowercase letter, a number, and a special character.
					MinLength:                 pulumi.Int(8),
				},
				// TODO: investigate this
				// MaintenanceWindow:
			},
			DeletionProtection: pulumi.Bool(isProdEnv(env)), // TODO: set to "true" in production
		},
			pulumi.IgnoreChanges([]string{"settings.diskSize"}), // avoid disk autoresize causing replacement
		)
		if err != nil {
			return err
		}
		// User for the database.
		dbAdminPW, err := random.NewRandomPassword(ctx, "db-admin-password", &random.RandomPasswordArgs{
			Length:  pulumi.Int(32),
			Special: pulumi.Bool(true),
			Lower:   pulumi.Bool(true),
			Upper:   pulumi.Bool(true),
			Numeric: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		dbAdmin, err := sql.NewUser(ctx, "db-admin-user", &sql.UserArgs{
			Name:     pulumi.String("pulumi_admin"),
			Instance: instance.Name,
			Password: dbAdminPW.Result,
		})
		if err != nil {
			return err
		}

		// IAM user
		dbDevOpsSA, err := sql.NewUser(ctx, "db-iam-devops", &sql.UserArgs{
			Name: devopsSA.Email.ApplyT(func(email string) *string {
				trimmed := strings.TrimSuffix(email, ".gserviceaccount.com")
				return &trimmed
			}).(pulumi.StringPtrOutput),
			Instance: instance.Name,
			Type:     pulumi.String("CLOUD_IAM_SERVICE_ACCOUNT"),
		},
			pulumi.DependsOn([]pulumi.Resource{instance, devopsSA}),
		)
		if err != nil {
			return err
		}
		dbApiSrvcSA, err := sql.NewUser(ctx, "db-iam-api-service", &sql.UserArgs{
			Name: apiSrvcSA.Email.ApplyT(func(email string) *string {
				trimmed := strings.TrimSuffix(email, ".gserviceaccount.com")
				return &trimmed
			}).(pulumi.StringPtrOutput),
			Instance: instance.Name,
			Type:     pulumi.String("CLOUD_IAM_SERVICE_ACCOUNT"),
		},
			pulumi.DependsOn([]pulumi.Resource{instance, apiSrvcSA}),
		)
		if err != nil {
			return err
		}
		dbAdminEmailIAM, err := sql.NewUser(ctx, "db-iam-admin-email", &sql.UserArgs{
			Name:     pulumi.String(adminEmail),
			Instance: instance.Name,
			Type:     pulumi.String("CLOUD_IAM_USER"),
		},
			pulumi.DependsOn([]pulumi.Resource{instance}),
		)
		if err != nil {
			return err
		}

		// Create a database inside the instance.
		db, err := sql.NewDatabase(ctx, "db-main", &sql.DatabaseArgs{
			Name:     pulumi.String(DB_STORAGE_ABBR),
			Instance: instance.Name,
		})
		if err != nil {
			return err
		}

		// Provider for the database.
		pgProvider, err := postgresql.NewProvider(ctx, "pg-provider", &postgresql.ProviderArgs{
			Database: db.Name,
			// Host:      instance.PublicIpAddress,
			// Port:      pulumi.Int(5432),

			// TODO: use these when using Cloud SQL Auth Proxy
			Host:    pulumi.String("127.0.0.1"),
			Port:    pulumi.Int(26257),
			Sslmode: pulumi.String("disable"),

			Username:  dbAdmin.Name,
			Password:  dbAdmin.Password,
			Superuser: pulumi.Bool(true),
		},
			pulumi.DependsOn([]pulumi.Resource{db}),
		)
		if err != nil {
			return err
		}

		// Enable the extensions.
		pgExts := []string{"citext", "tablefunc", "btree_gist", "pgaudit", "uuid-ossp"}
		for _, ext := range pgExts {
			_, err = postgresql.NewExtension(ctx, fmt.Sprintf("pg-ext-%s", ext), &postgresql.ExtensionArgs{
				Database: db.Name,
				Name:     pulumi.String(ext),
			},
				pulumi.Provider(pgProvider),
			)
			if err != nil {
				return err
			}
		}

		// Admin role is the owner of the database, which has all the privileges.
		dbRolAdmin, err := postgresql.NewRole(ctx, "pg-role-admin", &postgresql.RoleArgs{
			Name: pulumi.String(DB_STORAGE_ABBR + "_admin"),
		},
			pulumi.Provider(pgProvider),
		)
		if err != nil {
			return err
		}
		// Change the owner of the schema to the admin role.
		dbSchema, err := postgresql.NewSchema(ctx, "pg-schema-public", &postgresql.SchemaArgs{
			Database:    db.Name,
			Name:        pulumi.String("public"),
			Owner:       dbRolAdmin.Name,
			DropCascade: pulumi.Bool(!isProdEnv(env)),
		},
			pulumi.Provider(pgProvider),
		)
		if err != nil {
			return err
		}
		// Grant admin role to the service account.
		dbGrantedAdmins := []struct {
			name string
			role pulumi.StringOutput
		}{
			{
				name: "devops-sa",
				role: dbDevOpsSA.Name,
			},
			{
				name: "admin-email-iam",
				role: dbAdminEmailIAM.Name,
			},
		}
		for _, dbGrantedAdmin := range dbGrantedAdmins {
			grantedAdmin, err := postgresql.NewGrantRole(ctx, fmt.Sprintf("pg-grant-%s-role-admin", dbGrantedAdmin.name), &postgresql.GrantRoleArgs{
				Role:      dbGrantedAdmin.role,
				GrantRole: dbRolAdmin.Name,
			},
				pulumi.Provider(pgProvider),
			)
			if err != nil {
				return err
			}
			_, err = local.NewCommand(ctx, fmt.Sprintf("pg-%s-set-role-admin", dbGrantedAdmin.name), &local.CommandArgs{
				Create: pulumi.String(`PGSSLMODE=$SSLMODE PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USER -d $NAME -c "ALTER ROLE \"${ROLE}\" SET ROLE \"${SETROLE}\";"`),
				Environment: pulumi.StringMap{
					"HOST":     pulumi.String("127.0.0.1"),
					"PORT":     pulumi.String("26257"),
					"SSLMODE":  pulumi.String("disable"),
					"USER":     dbAdmin.Name,
					"NAME":     db.Name,
					"PASSWORD": dbAdminPW.Result,
					"ROLE":     dbGrantedAdmin.role,
					"SETROLE":  dbRolAdmin.Name,
				},
			},
				pulumi.DependsOn([]pulumi.Resource{grantedAdmin}),
			)
			if err != nil {
				return err
			}
		}

		// User role is the user of the database, which doesn't have all the privileges.
		dbRolUsr, err := postgresql.NewRole(ctx, "pg-role-user", &postgresql.RoleArgs{
			Name: pulumi.String(DB_STORAGE_ABBR + "_user"),
		},
			pulumi.DependsOn([]pulumi.Resource{instance}),
			pulumi.Provider(pgProvider),
		)
		if err != nil {
			return err
		}
		// Grant usage to the role.
		_, err = postgresql.NewGrant(ctx, "pg-grant-database-user", &postgresql.GrantArgs{
			Database:   db.Name,
			Role:       dbRolUsr.Name,
			ObjectType: pulumi.String("database"),
			Privileges: pulumi.StringArray{
				pulumi.String("CONNECT"),
			},
		},
			pulumi.Provider(pgProvider),
		)
		if err != nil {
			return err
		}
		pgObjPrivs := []struct {
			objType string
			privs   []string
		}{
			{"schema", []string{"USAGE", "CREATE"}},
			{"table", []string{"SELECT", "INSERT", "UPDATE", "DELETE", "TRUNCATE", "REFERENCES", "TRIGGER"}},
			{"sequence", []string{"USAGE", "SELECT", "UPDATE"}},
			{"function", []string{"EXECUTE"}},
			{"procedure", []string{"EXECUTE"}},
			{"routine", []string{"EXECUTE"}},
		}
		for _, objPrivs := range pgObjPrivs {
			privs := make(pulumi.StringArray, len(objPrivs.privs))
			for i, priv := range objPrivs.privs {
				privs[i] = pulumi.String(priv)
			}
			_, err = postgresql.NewGrant(ctx, fmt.Sprintf("pg-grant-%s-user", objPrivs.objType), &postgresql.GrantArgs{
				Database:   db.Name,
				Schema:     pulumi.String("public"),
				Role:       dbRolUsr.Name,
				ObjectType: pulumi.String(objPrivs.objType),
				Objects:    pulumi.StringArray{}, // empty means all objects
				Privileges: privs,
			},
				pulumi.Provider(pgProvider),
				pulumi.DependsOn([]pulumi.Resource{dbSchema}),
			)
			if err != nil {
				return err
			}
		}
		pgDefaultPrivs := []struct {
			objType string
			privs   []string
		}{
			{"schema", []string{"USAGE", "CREATE"}},
			{"table", []string{"SELECT", "INSERT", "UPDATE", "DELETE", "TRUNCATE", "REFERENCES", "TRIGGER"}},
			{"sequence", []string{"USAGE", "SELECT", "UPDATE"}},
			{"function", []string{"EXECUTE"}},
			{"type", []string{"USAGE"}},
		}
		for _, objPrivs := range pgDefaultPrivs {
			privs := make(pulumi.StringArray, len(objPrivs.privs))
			for i, priv := range objPrivs.privs {
				privs[i] = pulumi.String(priv)
			}
			var targetSchema pulumi.String
			if objPrivs.objType != "schema" {
				// Can't specify target schema when setting default privileges for "schema" object type.
				targetSchema = "public"
			}
			_, err = postgresql.NewDefaultPrivileges(ctx, fmt.Sprintf("pg-default-privs-%s-user", objPrivs.objType), &postgresql.DefaultPrivilegesArgs{
				Database:   db.Name,
				Schema:     targetSchema,
				Owner:      dbRolAdmin.Name, // Any object created by the admin role should be accessible by the user role.
				Role:       dbRolUsr.Name,
				ObjectType: pulumi.String(objPrivs.objType),
				Privileges: privs,
			},
				pulumi.Provider(pgProvider),
				pulumi.DependsOn([]pulumi.Resource{dbSchema}),
			)
			if err != nil {
				return err
			}
		}
		// Ensures that the db account is set to the correct role.
		dbGrantedUsrs := []struct {
			name string
			role pulumi.StringOutput
		}{
			{
				name: "api-service-sa",
				role: dbApiSrvcSA.Name,
			},
		}
		for _, dbGrantedUsr := range dbGrantedUsrs {
			grantedUsr, err := postgresql.NewGrantRole(ctx, fmt.Sprintf("pg-grant-%s-role-user", dbGrantedUsr.name), &postgresql.GrantRoleArgs{
				Role:      dbGrantedUsr.role,
				GrantRole: dbRolUsr.Name,
			},
				pulumi.Provider(pgProvider),
			)
			if err != nil {
				return err
			}
			_, err = local.NewCommand(ctx, fmt.Sprintf("pg-%s-set-role-user", dbGrantedUsr.name), &local.CommandArgs{
				Create: pulumi.String(`PGSSLMODE=$SSLMODE PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USER -d $NAME -c "ALTER ROLE \"${ROLE}\" SET ROLE \"${SETROLE}\";"`),
				Environment: pulumi.StringMap{
					"HOST":     pulumi.String("127.0.0.1"),
					"PORT":     pulumi.String("26257"),
					"SSLMODE":  pulumi.String("disable"),
					"USER":     dbAdmin.Name,
					"NAME":     db.Name,
					"PASSWORD": dbAdminPW.Result,
					"ROLE":     dbGrantedUsr.role,
					"SETROLE":  dbRolUsr.Name,
				},
			},
				pulumi.DependsOn([]pulumi.Resource{grantedUsr}),
			)
			if err != nil {
				return err
			}
		}

		// -------------------------------------------------------------------------------------------------------------
		// Cloud Run Service

		webService, err := cloudrunv2.NewService(ctx, "web-service", &cloudrunv2.ServiceArgs{
			Name:               pulumi.String("web-service"),
			Location:           pulumi.String("asia-east1"),
			DeletionProtection: pulumi.Bool(isProdEnv(env)),
			Ingress:            pulumi.String("INGRESS_TRAFFIC_ALL"),
			Template: &cloudrunv2.ServiceTemplateArgs{
				ServiceAccount:                webSrvcSA.Email,
				ExecutionEnvironment:          pulumi.String("EXECUTION_ENVIRONMENT_GEN1"),
				Timeout:                       pulumi.String("5s"),
				MaxInstanceRequestConcurrency: pulumi.Int(80),
				SessionAffinity:               pulumi.Bool(true),
				Scaling: &cloudrunv2.ServiceTemplateScalingArgs{
					MinInstanceCount: pulumi.Int(0),
					MaxInstanceCount: pulumi.Int(2),
				},
				Containers: cloudrunv2.ServiceTemplateContainerArray{
					&cloudrunv2.ServiceTemplateContainerArgs{
						Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
						Resources: &cloudrunv2.ServiceTemplateContainerResourcesArgs{
							CpuIdle: pulumi.Bool(true),
							Limits: pulumi.StringMap{
								"cpu":    pulumi.String("1"),
								"memory": pulumi.String("128Mi"),
							},
							StartupCpuBoost: pulumi.Bool(true),
						},
					},
				},
			},
			Traffics: cloudrunv2.ServiceTrafficArray{
				&cloudrunv2.ServiceTrafficArgs{
					Type:    pulumi.String("TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"),
					Percent: pulumi.Int(100),
				},
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.IgnoreChanges([]string{"template.containers[0].image"}), // ignore changes to the image
		)
		if err != nil {
			return err
		}
		_, err = cloudrunv2.NewServiceIamBinding(ctx, "web-service-iam-binding", &cloudrunv2.ServiceIamBindingArgs{
			Name:     webService.Name,
			Location: webService.Location,
			Role:     pulumi.String("roles/run.invoker"),
			Members: pulumi.StringArray{
				pulumi.String("allUsers"),
			},
		})
		if err != nil {
			return err
		}
		webServiceDomain, err := cloudrun.NewDomainMapping(ctx, "web-service-domain", &cloudrun.DomainMappingArgs{
			Name:     pulumi.String(domainName(env, "")),
			Location: webService.Location,
			Metadata: &cloudrun.DomainMappingMetadataArgs{
				Namespace: pulumi.String(gcpProject),
			},
			Spec: &cloudrun.DomainMappingSpecArgs{
				RouteName: webService.Name,
			},
		})
		if err != nil {
			return err
		}

		apiService, err := cloudrunv2.NewService(ctx, "api-service", &cloudrunv2.ServiceArgs{
			Name:               pulumi.String("api-service"),
			Location:           pulumi.String("asia-east1"),
			DeletionProtection: pulumi.Bool(isProdEnv(env)),
			Ingress:            pulumi.String("INGRESS_TRAFFIC_ALL"),
			Template: &cloudrunv2.ServiceTemplateArgs{
				ServiceAccount:                apiSrvcSA.Email,
				ExecutionEnvironment:          pulumi.String("EXECUTION_ENVIRONMENT_GEN1"),
				Timeout:                       pulumi.String("10s"),
				MaxInstanceRequestConcurrency: pulumi.Int(80),
				SessionAffinity:               pulumi.Bool(true),
				Scaling: &cloudrunv2.ServiceTemplateScalingArgs{
					MinInstanceCount: pulumi.Int(0),
					MaxInstanceCount: pulumi.Int(5),
				},
				Containers: cloudrunv2.ServiceTemplateContainerArray{
					&cloudrunv2.ServiceTemplateContainerArgs{
						Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
						Resources: &cloudrunv2.ServiceTemplateContainerResourcesArgs{
							CpuIdle: pulumi.Bool(true),
							Limits: pulumi.StringMap{
								"cpu":    pulumi.String("1"),
								"memory": pulumi.String("256Mi"),
							},
							StartupCpuBoost: pulumi.Bool(true),
						},
						Ports: &cloudrunv2.ServiceTemplateContainerPortsArgs{
							ContainerPort: pulumi.Int(3000),
							Name:          pulumi.String("http1"),
						},
						Envs: cloudrunv2.ServiceTemplateContainerEnvArray{
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name: pulumi.String("WEB_CORS_ALLOWED_ORIGINS"),
								Value: webServiceDomain.Name.ApplyT(func(domain string) string {
									return fmt.Sprintf("https://%s", domain)
								}).(pulumi.StringOutput),
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name:  pulumi.String("DB_CLOUD_SQL_CONNECTION_NAME"),
								Value: instance.ConnectionName,
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name:  pulumi.String("DB_USER"),
								Value: dbApiSrvcSA.Name,
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name:  pulumi.String("DB_NAME"),
								Value: db.Name,
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name: pulumi.String("SESSION_SECRET_KEY"),
								ValueSource: &cloudrunv2.ServiceTemplateContainerEnvValueSourceArgs{
									SecretKeyRef: &cloudrunv2.ServiceTemplateContainerEnvValueSourceSecretKeyRefArgs{
										Secret:  sessionSecret.Name,
										Version: pulumi.String("latest"),
									},
								},
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name:  pulumi.String("SESSION_SAME_SITE"),
								Value: pulumi.String("None"), // TODO: set to "Lax" in production where we have a domain that supports it
							},
							&cloudrunv2.ServiceTemplateContainerEnvArgs{
								Name:  pulumi.String("STORAGE_BUCKET"),
								Value: bucket.Name,
							},
						},
						//	TODO: StartupProbe / LivenessProbe
					},
				},
			},
			Traffics: cloudrunv2.ServiceTrafficArray{
				&cloudrunv2.ServiceTrafficArgs{
					Type:    pulumi.String("TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"),
					Percent: pulumi.Int(100),
				},
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.IgnoreChanges([]string{"template.containers[0].image"}), // ignore changes to the image
		)
		if err != nil {
			return err
		}
		_, err = cloudrunv2.NewServiceIamBinding(ctx, "api-service-iam-binding", &cloudrunv2.ServiceIamBindingArgs{
			Name:     apiService.Name,
			Location: apiService.Location,
			Role:     pulumi.String("roles/run.invoker"),
			Members: pulumi.StringArray{
				pulumi.String("allUsers"),
			},
		})
		if err != nil {
			return err
		}
		apiServiceDomain, err := cloudrun.NewDomainMapping(ctx, "api-service-domain", &cloudrun.DomainMappingArgs{
			Name:     pulumi.String(domainName(env, "api")),
			Location: apiService.Location,
			Metadata: &cloudrun.DomainMappingMetadataArgs{
				Namespace: pulumi.String(gcpProject),
			},
			Spec: &cloudrun.DomainMappingSpecArgs{
				RouteName: apiService.Name,
			},
		})
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Cloud Run Jobs

		migrationJob, err := cloudrunv2.NewJob(ctx, "migration-job", &cloudrunv2.JobArgs{
			Name:               pulumi.String("migration-job"),
			Location:           pulumi.String("asia-east1"),
			DeletionProtection: pulumi.Bool(isProdEnv(env)),
			Template: &cloudrunv2.JobTemplateArgs{
				TaskCount: pulumi.Int(1),
				Template: &cloudrunv2.JobTemplateTemplateArgs{
					ServiceAccount:       devopsSA.Email,
					MaxRetries:           pulumi.Int(0),
					Timeout:              pulumi.String("120s"),
					ExecutionEnvironment: pulumi.String("EXECUTION_ENVIRONMENT_GEN2"),
					Containers: cloudrunv2.JobTemplateTemplateContainerArray{
						&cloudrunv2.JobTemplateTemplateContainerArgs{
							Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/job"),
							Resources: &cloudrunv2.JobTemplateTemplateContainerResourcesArgs{
								Limits: pulumi.StringMap{
									"cpu":    pulumi.String("1"),
									"memory": pulumi.String("512Mi"),
								},
							},
							Envs: cloudrunv2.JobTemplateTemplateContainerEnvArray{
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_CLOUD_SQL_CONNECTION_NAME"),
									Value: instance.ConnectionName,
								},
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_USER"),
									Value: dbDevOpsSA.Name,
								},
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_NAME"),
									Value: db.Name,
								},
							},
						},
					},
				},
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.IgnoreChanges([]string{"template.template.containers[0].image"}), // ignore changes to the image
		)
		if err != nil {
			return err
		}

		workerJob, err := cloudrunv2.NewJob(ctx, "worker-job", &cloudrunv2.JobArgs{
			Name:               pulumi.String("worker-job"),
			Location:           pulumi.String("asia-east1"),
			DeletionProtection: pulumi.Bool(isProdEnv(env)),
			Template: &cloudrunv2.JobTemplateArgs{
				TaskCount: pulumi.Int(1),
				Template: &cloudrunv2.JobTemplateTemplateArgs{
					ServiceAccount:       devopsSA.Email,
					MaxRetries:           pulumi.Int(0),
					Timeout:              pulumi.String("120s"),
					ExecutionEnvironment: pulumi.String("EXECUTION_ENVIRONMENT_GEN2"),
					Containers: cloudrunv2.JobTemplateTemplateContainerArray{
						&cloudrunv2.JobTemplateTemplateContainerArgs{
							Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/job"),
							Resources: &cloudrunv2.JobTemplateTemplateContainerResourcesArgs{
								Limits: pulumi.StringMap{
									"cpu":    pulumi.String("1"),
									"memory": pulumi.String("512Mi"),
								},
							},
							Envs: cloudrunv2.JobTemplateTemplateContainerEnvArray{
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_CLOUD_SQL_CONNECTION_NAME"),
									Value: instance.ConnectionName,
								},
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_USER"),
									Value: dbDevOpsSA.Name,
								},
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("DB_NAME"),
									Value: db.Name,
								},
								&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
									Name:  pulumi.String("STORAGE_BUCKET"),
									Value: bucket.Name,
								},
							},
						},
					},
				},
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.IgnoreChanges([]string{"template.template.containers[0].image"}), // ignore changes to the image
		)
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Cloud Scheduler

		// jobs := []struct {
		// 	name     string
		// 	schedule string
		// }{
		// 	{jobname.JobPaidLeaveDailyTrigger, "1 0 * * *"},           // 00:01 at Asia/Taipei time
		// 	{jobname.JobPaidLeaveDailyCleanUp, "11 0 * * *"},          // 00:11 at Asia/Taipei time
		// 	{jobname.JobEmployeeSalaryFormDailyTrigger, "21 0 * * *"}, // 00:21 at Asia/Taipei time
		// 	{jobname.JobExportSalaryExcelsDailyTrigger, "31 0 * * *"}, // 00:31 at Asia/Taipei time
		// }
		// for _, j := range jobs {
		// 	invokeBase64encode, err := std.Base64encode(ctx, &std.Base64encodeArgs{
		// 		Input: fmt.Sprintf(`{"overrides": {"containerOverrides": [{"args": [], "env": [{"name": "JOB_NAME", "value": "%s"}]}], "taskCount": 1, "timeout": "300s" }}`, j.name),
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}

		// 	_, err = cloudscheduler.NewJob(ctx, fmt.Sprintf("scheduler-%s", j.name), &cloudscheduler.JobArgs{
		// 		Region:          pulumi.String("asia-east1"),
		// 		Name:            pulumi.String(j.name),
		// 		Schedule:        pulumi.String(j.schedule),
		// 		TimeZone:        pulumi.String("Asia/Taipei"),
		// 		AttemptDeadline: pulumi.String("300s"),
		// 		RetryConfig: &cloudscheduler.JobRetryConfigArgs{
		// 			RetryCount: pulumi.Int(1),
		// 		},
		// 		HttpTarget: &cloudscheduler.JobHttpTargetArgs{
		// 			HttpMethod: pulumi.String("POST"),
		// 			Uri: workerJob.Name.ApplyT(func(name string) string {
		// 				return fmt.Sprintf("https://run.googleapis.com/v2/projects/%s/locations/%s/jobs/%s:run", gcpProject, "asia-east1", name)
		// 			}).(pulumi.StringOutput),
		// 			OauthToken: &cloudscheduler.JobHttpTargetOauthTokenArgs{
		// 				ServiceAccountEmail: devopsSA.Email,
		// 			},
		// 			Headers: pulumi.StringMap{
		// 				"Content-Type": pulumi.String("application/json"),
		// 			},
		// 			Body: pulumi.String(invokeBase64encode.Result),
		// 		},
		// 	},
		// 		pulumi.DependsOn(enabledService),
		// 		pulumi.DependsOn([]pulumi.Resource{workerJob}),
		// 	)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// -------------------------------------------------------------------------------------------------------------
		// Cloud Build

		_, err = cloudbuild.NewTrigger(ctx, "web-service-trigger", &cloudbuild.TriggerArgs{
			Location:       pulumi.String("asia-east1"),
			Filename:       pulumi.String("deploy/cloudbuild.admin.yaml"),
			ServiceAccount: devopsSA.Name,
			Github: &cloudbuild.TriggerGithubArgs{
				Owner: pulumi.String(REPO_OWNER),
				Name:  pulumi.String(REPO_NAME),
				Push: &cloudbuild.TriggerGithubPushArgs{
					Branch: pulumi.String(fmt.Sprintf("^%s$", gitBranch(env))),
				},
			},
			IncludeBuildLogs: pulumi.String("INCLUDE_BUILD_LOGS_WITH_STATUS"),
			Substitutions: pulumi.StringMap{
				"_AR_HOSTNAME": ar.Location.ApplyT(func(loc string) string {
					return fmt.Sprintf("%s-docker.pkg.dev", loc)
				}).(pulumi.StringOutput),
				"_AR_REPO_ID":    ar.RepositoryId,
				"_SERVICE_NAME":  webService.Name,
				"_DEPLOY_REGION": webService.Location,
				"_ENV":           pulumi.String(webServiceEnv(env)),
			},
			Tags: pulumi.StringArray{
				pulumi.String("web-service"),
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.DependsOn([]pulumi.Resource{webService}),
		)
		if err != nil {
			return err
		}

		_, err = cloudbuild.NewTrigger(ctx, "api-service-trigger", &cloudbuild.TriggerArgs{
			Location:       pulumi.String("asia-east1"),
			Filename:       pulumi.String("deploy/cloudbuild.api.yaml"),
			ServiceAccount: devopsSA.Name,
			Github: &cloudbuild.TriggerGithubArgs{
				Owner: pulumi.String(REPO_OWNER),
				Name:  pulumi.String(REPO_NAME),
				Push: &cloudbuild.TriggerGithubPushArgs{
					Branch: pulumi.String(fmt.Sprintf("^%s$", gitBranch(env))),
				},
			},
			IncludeBuildLogs: pulumi.String("INCLUDE_BUILD_LOGS_WITH_STATUS"),
			Substitutions: pulumi.StringMap{
				"_AR_HOSTNAME": ar.Location.ApplyT(func(loc string) string {
					return fmt.Sprintf("%s-docker.pkg.dev", loc)
				}).(pulumi.StringOutput),
				"_AR_REPO_ID":         ar.RepositoryId,
				"_SERVICE_NAME":       apiService.Name,
				"_MIGRATION_JOB_NAME": migrationJob.Name,
				"_DEPLOY_REGION":      apiService.Location,
			},
			Tags: pulumi.StringArray{
				pulumi.String("api-service"),
			},
		},
			pulumi.DependsOn(enabledService),
			pulumi.DependsOn([]pulumi.Resource{apiService}),
		)
		if err != nil {
			return err
		}

		_, err = cloudbuild.NewTrigger(ctx, "worker-service-trigger", &cloudbuild.TriggerArgs{
			Location:       pulumi.String("asia-east1"),
			Filename:       pulumi.String("deploy/cloudbuild.worker.yaml"),
			ServiceAccount: devopsSA.Name,
			Github: &cloudbuild.TriggerGithubArgs{
				Owner: pulumi.String(REPO_OWNER),
				Name:  pulumi.String(REPO_NAME),
				Push: &cloudbuild.TriggerGithubPushArgs{
					Branch: pulumi.String(fmt.Sprintf("^%s$", gitBranch(env))),
				},
			},
			IncludeBuildLogs: pulumi.String("INCLUDE_BUILD_LOGS_WITH_STATUS"),
			Substitutions: pulumi.StringMap{
				"_AR_HOSTNAME": ar.Location.ApplyT(func(loc string) string {
					return fmt.Sprintf("%s-docker.pkg.dev", loc)
				}).(pulumi.StringOutput),
				"_AR_REPO_ID":    ar.RepositoryId,
				"_SERVICE_NAME":  workerJob.Name,
				"_DEPLOY_REGION": workerJob.Location,
			},
			Tags: pulumi.StringArray{
				pulumi.String("worker-service"),
			},
		},
			pulumi.DependsOn(enabledService),
		)
		if err != nil {
			return err
		}

		// -------------------------------------------------------------------------------------------------------------
		// Outputs

		ctx.Export("GCS bucketName", bucket.Url)
		ctx.Export("SQL instanceName", instance.Name)
		ctx.Export("SQL databaseName", db.Name)
		ctx.Export("SQL admin username", dbAdmin.Name)
		ctx.Export("Cloud Run web service URL", webService.Uri)
		ctx.Export("Cloud Run web service domain", webServiceDomain.Name)
		ctx.Export("Cloud Run api service URL", apiService.Uri)
		ctx.Export("Cloud Run api service domain", apiServiceDomain.Name)

		return nil
	})
}

func isProdEnv(env string) bool {
	return env == "prod"
}

func domainName(env, subdomain string) string {
	const defaultDomain = DEFAULT_DOMAIN

	if isProdEnv(env) {
		if subdomain == "" {
			return defaultDomain
		}
		return fmt.Sprintf("%s.%s", subdomain, defaultDomain)
	}

	if subdomain == "" {
		return fmt.Sprintf("%s.%s", env, defaultDomain)
	}
	return fmt.Sprintf("%s-%s.%s", subdomain, env, defaultDomain)
}

func gitBranch(env string) string {
	if isProdEnv(env) {
		return "prod"
	}
	return "main"
}

func webServiceEnv(env string) string {
	if isProdEnv(env) {
		return "production"
	}
	return "staging"
}
