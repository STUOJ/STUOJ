package judge0

type JudgeSubmission struct {
	SourceCode     string  `json:"source_code"`
	LanguageId     uint64  `json:"language_id"`
	Stdin          string  `json:"stdin,omitempty"`
	ExpectedOutput string  `json:"expected_output,omitempty"`
	CPUTimeLimit   float64 `json:"cpu_time_limit,omitempty"`
	MemoryLimit    uint64  `json:"memory_limit,omitempty"`
}

type JudgeResult struct {
	Stdout        string            `json:"stdout,omitempty"`
	Time          string            `json:"time,omitempty"`
	Memory        float64           `json:"memory,omitempty"`
	Stderr        string            `json:"stderr,omitempty"`
	Token         string            `json:"token"`
	CompileOutput string            `json:"compile_output,omitempty"`
	Message       string            `json:"message,omitempty"`
	Status        JudgeResultStatus `json:"status"`
}

type JudgeResultStatus struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
}

type JudgeSubmissionResults struct {
	Submissions []JudgeResult              `json:"submissions"`
	Meta        JudgeSubmissionResultsMeta `json:"meta"`
}

type JudgeSubmissionResultsMeta struct {
	CurrentPage uint64 `json:"current_page"`
	NextPage    uint64 `json:"next_page"`
	PrevPage    uint64 `json:"prev_page"`
	TotalPages  uint64 `json:"total_pages"`
	TotalCount  uint64 `json:"total_count"`
}

type JudgeConfigInfo struct {
	MaintenanceMode                           bool     `json:"maintenance_mode"`
	EnableWaitResult                          bool     `json:"enable_wait_result"`
	EnableCompilerOptions                     bool     `json:"enable_compiler_options"`
	AllowedLanguagesForCompileOptions         []string `json:"allowed_languages_for_compile_options"`
	EnableCommandLineArguments                bool     `json:"enable_command_line_arguments"`
	EnableSubmissionDelete                    bool     `json:"enable_submission_delete"`
	EnableCallbacks                           bool     `json:"enable_callbacks"`
	CallbacksMaxTries                         uint64   `json:"callbacks_max_tries"`
	CallbacksTimeout                          float64  `json:"callbacks_timeout"`
	EnableAdditionalFiles                     bool     `json:"enable_additional_files"`
	MaxQueueSize                              uint64   `json:"max_queue_size"`
	CpuTimeLimit                              float64  `json:"cpu_time_limit"`
	MaxCpuTimeLimit                           float64  `json:"max_cpu_time_limit"`
	CpuExtraTime                              float64  `json:"cpu_extra_time"`
	MaxCpuExtraTime                           float64  `json:"max_cpu_extra_time"`
	WallTimeLimit                             float64  `json:"wall_time_limit"`
	MaxWallTimeLimit                          float64  `json:"max_wall_time_limit"`
	MemoryLimit                               uint64   `json:"memory_limit"`
	MaxMemoryLimit                            uint64   `json:"max_memory_limit"`
	StackLimit                                uint64   `json:"stack_limit"`
	MaxStackLimit                             uint64   `json:"max_stack_limit"`
	MaxProcessesAndOrThreads                  uint64   `json:"max_processes_and_or_threads"`
	MaxMaxProcessesAndOrThreads               uint64   `json:"max_max_processes_and_or_threads"`
	EnablePerProcessAndThreadTimeLimit        bool     `json:"enable_per_process_and_thread_time_limit"`
	AllowEnablePerProcessAndThreadTimeLimit   bool     `json:"allow_enable_per_process_and_thread_time_limit"`
	EnablePerProcessAndThreadMemoryLimit      bool     `json:"enable_per_process_and_thread_memory_limit"`
	AllowEnablePerProcessAndThreadMemoryLimit bool     `json:"allow_enable_per_process_and_thread_memory_limit"`
	MaxFileSize                               uint64   `json:"max_file_size"`
	MaxMaxFileSize                            uint64   `json:"max_max_file_size"`
	NumberOfRuns                              uint64   `json:"number_of_runs"`
	MaxNumberOfRuns                           uint64   `json:"max_number_of_runs"`
	RedirectStderrToStdout                    bool     `json:"redirect_stderr_to_stdout"`
	MaxExtractSize                            uint64   `json:"max_extract_size"`
	EnableBatchedSubmissions                  bool     `json:"enable_batched_submissions"`
	MaxSubmissionBatchSize                    uint64   `json:"max_submission_batch_size"`
	SubmissionCacheDuration                   float64  `json:"submission_cache_duration"`
	UseDocsAsHomepage                         bool     `json:"use_docs_as_homepage"`
	AllowEnableNetwork                        bool     `json:"allow_enable_network"`
	EnableNetwork                             bool     `json:"enable_network"`
}

type JudgeSystemInfo struct {
	Architecture       string `json:"Architecture"`
	CPUOpModeS         string `json:"CPU op-mode(s)"`
	ByteOrder          string `json:"Byte Order"`
	AddressSizes       string `json:"Address sizes"`
	CPUS               string `json:"CPU(s)"`
	OnLineCPUSList     string `json:"On-line CPU(s) list"`
	ThreadSPerCore     string `json:"Thread(s) per core"`
	CoreSPerSocket     string `json:"Core(s) per socket"`
	SocketS            string `json:"Socket(s)"`
	VendorID           string `json:"Vendor ID"`
	CPUFamily          string `json:"CPU family"`
	Model              string `json:"Model"`
	ModelName          string `json:"Model name"`
	Stepping           string `json:"Stepping"`
	CPUMHz             string `json:"CPU MHz"`
	BogoMIPS           string `json:"BogoMIPS"`
	Virtualization     string `json:"Virtualization"`
	HypervisorVendor   string `json:"Hypervisor vendor"`
	VirtualizationType string `json:"Virtualization type"`
	L1dCache           string `json:"L1d cache"`
	L1iCache           string `json:"L1i cache"`
	L2Cache            string `json:"L2 cache"`
	L3Cache            string `json:"L3 cache"`
	Flags              string `json:"Flags"`
	Mem                string `json:"Mem"`
	Swap               string `json:"Swap"`
}

type JudgeStatistics struct {
	CreatedAt   string                     `json:"created_at"`
	CachedUntil string                     `json:"cached_until"`
	Submissions JudgeSubmissionsStatistics `json:"submissions"`
	Languages   []JudgeLanguageStatistics  `json:"languages"`
	Statuses    []JudgeStatusStatistics    `json:"statuses"`
	Database    JudgeDatabaseStatistics    `json:"db"`
}

type JudgeSubmissionsStatistics struct {
	Total      uint64                 `json:"total"`
	Today      uint64                 `json:"today"`
	Last30Days map[string]interface{} `json:"last_30_days"`
}

type JudgeLanguageStatistics struct {
	Language JudgeLanguageInfo `json:"language"`
	Count    int32             `json:"count"`
}

type JudgeLanguageInfo struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type JudgeStatusStatistics struct {
	Status JudgeStatus `json:"status"`
	Count  uint64      `json:"count"`
}

type JudgeStatus struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type JudgeDatabaseStatistics struct {
	SizePretty  string `json:"size_pretty"`
	SizeInBytes uint64 `json:"size_in_bytes"`
}

type JudgeAbout struct {
	Version    string `json:"version"`
	Homepage   string `json:"homepage"`
	SourceCode string `json:"source_code"`
	Maintainer string `json:"maintainer"`
}

type JudgeWorker struct {
	Queue     string `json:"queue,omitempty"`
	Size      uint64 `json:"size,omitempty"`
	Available uint64 `json:"available,omitempty"`
	Idle      uint64 `json:"idle,omitempty"`
	Working   uint64 `json:"working,omitempty"`
	Paused    uint64 `json:"paused,omitempty"`
	Failed    uint64 `json:"failed,omitempty"`
}

type JudgeLanguage struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
