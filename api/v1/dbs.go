package v1

type HugePages struct {
	Enabled bool   `json:"enabled,omitempty"`
	Size    string `json:"size,omitempty"`
	Memory  string `json:"memory,omitempty"`
}

type Pg struct {
	Enabled        bool              `json:"enabled,omitempty"`
	ServiceAccount string            `json:"serviceAccount"`
	SecretName     string            `json:"secretName,omitempty"`
	Image          string            `json:"image,omitempty"`
	Port           int               `json:"port,omitempty"`
	StorageSize    string            `json:"storageSize,omitempty"`
	SvcName        string            `json:"svcName,omitempty"`
	StorageClass   string            `json:"storageClass,omitempty"`
	CPURequest     string            `json:"cpuRequest,omitempty"`
	MemoryRequest  string            `json:"memoryRequest,omitempty"`
	MaxConnections int               `json:"maxConnections,omitempty"`
	SharedBuffers  string            `json:"sharedBuffers,omitempty"`
	HugePages      HugePages         `json:"hugePages,omitempty"`
	Fixpg          bool              `json:"fixpg,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty"`
	Tolerations    map[string]string `json:"tolerations,omitempty"`
	CredsRef       string            `json:"credsRef"`
}

type Minio struct {
	Enabled        bool              `json:"enabled,omitempty"`
	ServiceAccount string            `json:"serviceAccount,omitempty"`
	Replicas       int               `json:"replicas,omitempty"`
	Image          string            `json:"image,omitempty"`
	Port           int               `json:"port,omitempty"`
	StorageSize    string            `json:"storageSize,omitempty"`
	SvcName        string            `json:"svcName,omitempty"`
	NodePort       int               `json:"nodePort,omitempty"`
	StorageClass   string            `json:"storageClass,omitempty"`
	CPURequest     string            `json:"cpuRequest,omitempty"`
	MemoryRequest  string            `json:"memoryRequest,omitempty"`
	SharedStorage  SharedStorage     `json:"sharedStorage,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty"`
	Tolerations    map[string]string `json:"tolerations,omitempty"`
}

type Redis struct {
	Enabled        bool              `json:"enabled,omitempty"`
	ServiceAccount string            `json:"serviceAccount"`
	Image          string            `json:"image,omitempty"`
	SvcName        string            `json:"svcName,omitempty"`
	Port           int               `json:"port,omitempty"`
	StorageSize    string            `json:"storageSize,omitempty"`
	StorageClass   string            `json:"storageClass,omitempty"`
	Limits         Limits            `json:"limits,omitempty"`
	Requests       Requests          `json:"requests,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty"`
	Tolerations    map[string]string `json:"tolerations,omitempty"`
	CredsRef       string            `json:"CredsRef"`
}

type Es struct {
	Enabled        bool              `json:"enabled,omitempty"`
	ServiceAccount string            `json:"serviceAccount,omitempty"`
	Image          string            `json:"image,omitempty"`
	Port           int               `json:"port,omitempty"`
	StorageSize    string            `json:"storageSize,omitempty"`
	SvcName        string            `json:"svcName,omitempty"`
	RunAsUser      int               `json:"runAsUser,omitempty"`
	FsGroup        int               `json:"fsGroup,omitempty"`
	NodePort       int               `json:"nodePort,omitempty"`
	StorageClass   string            `json:"storageClass,omitempty"`
	CPURequest     string            `json:"cpuRequest,omitempty"`
	MemoryRequest  string            `json:"memoryRequest,omitempty"`
	CPULimit       string            `json:"cpuLimit,omitempty"`
	MemoryLimit    string            `json:"memoryLimit,omitempty"`
	JavaOpts       string            `json:"javaOpts,omitempty"`
	PatchEsNodes   string            `json:"patchEsNodes,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty"`
	Tolerations    map[string]string `json:"tolerations,omitempty"`
	CredsRef       string            `json:"credsRef"`
}

type AppDbs struct {
	Pg    Pg    `json:"pg,omitempty"`
	Redis Redis `json:"redis,omitempty"`
	Minio Minio `json:"minio,omitempty"`
	Es    Es    `json:"es,omitempty"`
}

type InfraDbs struct {
	Redis Redis `json:"redis,omitempty"`
}

var minioDefaults = Minio{
	Enabled:        true,
	ServiceAccount: "minio",
	Replicas:       1,
	Image:          "docker.io/minio/minio:RELEASE.2020-09-17T04-49-20Z",
	Port:           9000,
	StorageSize:    "100Gi",
	SvcName:        "minio",
	NodePort:       30090,
	StorageClass:   "",
	CPURequest:     "1000m",
	MemoryRequest:  "2Gi",
	SharedStorage: SharedStorage{
		Enabled:          false,
		UseExistingClaim: "",
		ConsistentHash: ConsistentHash{
			Key:   "httpQueryParameterName",
			Value: "uploadId",
		},
	},
}

var pgDefault = Pg{
	Enabled:        true,
	ServiceAccount: "pg",
	SecretName:     "cnvrg-pg-secret",
	Image:          "centos/postgresql-12-centos7",
	Port:           5432,
	StorageSize:    "80Gi",
	SvcName:        "postgres",
	StorageClass:   "",
	CPURequest:     "4000m",
	MemoryRequest:  "4Gi",
	MaxConnections: 100,
	SharedBuffers:  "64MB",
	Fixpg:          true,
	NodeSelector:   nil,
	Tolerations:    nil,
	HugePages: HugePages{
		Enabled: false,
		Size:    "2Mi",
		Memory:  "",
	},
	CredsRef: "pg-creds",
}

var redisDefault = Redis{
	Enabled:        true,
	ServiceAccount: "redis",
	Image:          "docker.io/cnvrg/cnvrg-redis:v3.0.5.c2",
	SvcName:        "redis",
	Port:           6379,
	StorageSize:    "10Gi",
	StorageClass:   "",
	NodeSelector:   nil,
	Tolerations:    nil,
	CredsRef:       "redis-creds",
	Limits: Limits{
		CPU:    "1000m",
		Memory: "2Gi",
	},
	Requests: Requests{
		CPU:    "100m",
		Memory: "200Mi",
	},
}

var esDefault = Es{
	Enabled:        true,
	ServiceAccount: "es",
	Image:          "docker.io/cnvrg/cnvrg-es:v7.8.1.a1",
	Port:           9200,
	StorageSize:    "30Gi",
	SvcName:        "elasticsearch",
	RunAsUser:      1000,
	FsGroup:        1000,
	NodePort:       32200,
	StorageClass:   "",
	CPURequest:     "1000m",
	MemoryRequest:  "1Gi",
	CPULimit:       "2000m",
	MemoryLimit:    "4Gi",
	JavaOpts:       "",
	PatchEsNodes:   "true",
	NodeSelector:   nil,
	Tolerations:    nil,
	CredsRef:       "es-creds",
}

var appDbsDefaults = AppDbs{
	Pg:    pgDefault,
	Redis: redisDefault,
	Minio: minioDefaults,
	Es:    esDefault,
}

var infraDbsDefaults = InfraDbs{
	Redis: redisDefault,
}
