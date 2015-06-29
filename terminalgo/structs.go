package terminalgo


// DATA OUTPUT STRUCTS

type Terminal_List struct {
	Terminals []struct {
		CPU          string `json:"cpu"`
		RAM          string `json:"ram"`
		Diskspace    string `json:"diskspace"`
		Name         string `json:"name"`
		SnapshotID   string `json:"snapshot_id"`
		Temporary    string `json:"temporary"`
		CustomData   string `json:"custom_data"`
		Status       string `json:"status"`
		AllowSpot    string `json:"allow_spot"`
		ContainerKey string `json:"container_key"`
		Subdomain    string `json:"subdomain"`
		ContainerIP  string `json:"container_ip"`
		CreationTime string `json:"creation_time"`
		RequestID    string `json:"request_id"`
		InstanceType string `json:"instance_type"`
	} `json:"terminals"`
}

type Output_Snapshot struct {
	Snapshot struct {
		Title        string `json:"title"`
		Body         string `json:"body"`
		DisplayStyle string `json:"display_style"`
		Readme       string `json:"readme"`
		Pynb         string `json:"pynb"`
		Tags         string `json:"tags"`
		Category     string `json:"category"`
		Image        struct {
			Files  []string `json:"files"`
			Cdnuri string   `json:"cdnUri"`
		} `json:"image"`
		Diskspace  int    `json:"diskspace"`
		RAM        int    `json:"ram"`
		Public     bool   `json:"public"`
		StartCount int    `json:"start_count"`
		Createdat  string `json:"createdAt"`
		Comments   []struct {
			Createdat string `json:"createdAt"`
			Body      string `json:"body"`
			ID        string `json:"id"`
			User      struct {
				Username interface{} `json:"username"`
				Avatar   interface{} `json:"avatar"`
			} `json:"user"`
		} `json:"comments"`
		Featured    bool   `json:"featured"`
		Author      string `json:"author"`
		ID          string `json:"id"`
		Templatekey string `json:"templateKey"`
	} `json:"snapshot"`
}



type Output_User struct {
	User struct {
		Name         string `json:"name"`
		Username     string `json:"username"`
		URL          string `json:"url"`
		Company      string `json:"company"`
		Location     string `json:"location"`
		ProfileImage string `json:"profile_image"`
	} `json:"user"`
}

type Output_Snapshots_List struct {
	Snapshots []struct {
		Author       string  `json:"author"`
		Body         string  `json:"body"`
		Category     string  `json:"category"`
		CreatedAt    string  `json:"createdAt"`
		Diskspace    float64 `json:"diskspace"`
		DisplayStyle string  `json:"display_style"`
		Featured     bool    `json:"featured"`
		ID           string  `json:"id"`
		Image        struct {
			CdnUri string   `json:"cdnUri"`
			Files  []string `json:"files"`
		} `json:"image"`
		Public      bool    `json:"public"`
		Pynb        string  `json:"pynb"`
		Ram         float64 `json:"ram"`
		Readme      string  `json:"readme"`
		StartCount  float64 `json:"start_count"`
		Tags        string  `json:"tags"`
		TemplateKey string  `json:"templateKey"`
		Title       string  `json:"title"`
	} `json:"snapshots"`
}


type Output_Snapshot_Count struct {
	SnapshotCount int `json:"snapshot_count"`
}

type Output_Terminal struct {
	Terminal struct {
		AllowSpot    string `json:"allow_spot"`
		ContainerIp  string `json:"container_ip"`
		ContainerKey string `json:"container_key"`
		Cpu          string `json:"cpu"`
		CreationTime string `json:"creation_time"`
		CustomData   string `json:"custom_data"`
		Diskspace    string `json:"diskspace"`
		InstanceType string `json:"instance_type"`
		Name         string `json:"name"`
		Ram          string `json:"ram"`
		SnapshotID   string `json:"snapshot_id"`
		Status       string `json:"status"`
		Subdomain    string `json:"subdomain"`
		Temporary    string `json:"temporary"`
	} `json:"terminal"`
}

type Output_RequestID struct {
	RequestID string `json:"request_id"`
}

type Output_Req_Status struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
}

type Output_Status struct {
	Status    string `json:"status"`
}

type Output_Req_Success struct {
	RequestID string `json:"request_id"`
	Success    bool `json:"success"`
}


type Output_Snapshots_Priv_List struct {
	Snapshots []struct {
		Title string `json:"title"`
		Body string `json:"body"`
		DisplayStyle string `json:"display_style"`
		Readme string `json:"readme"`
		Pynb string `json:"pynb"`
		Tags string `json:"tags"`
		Category string `json:"category"`
		Image struct {
			Files []string `json:"files"`
			Cdnuri string `json:"cdnUri"`
		} `json:"image"`
		Diskspace int `json:"diskspace"`
		RAM int `json:"ram"`
		Public bool `json:"public"`
		StartCount int `json:"start_count"`
		Createdat string `json:"createdAt"`
		Parent string `json:"parent"`
		Children []interface{} `json:"children"`
		Featured bool `json:"featured"`
		Author string `json:"author"`
		ID string `json:"id"`
		Templatekey string `json:"templateKey"`
	} `json:"snapshots"`
}

type Output_Snapshot_Edit struct {
	Snapshot struct {
		Title string `json:"title"`
		Body string `json:"body"`
		DisplayStyle string `json:"display_style"`
		Readme string `json:"readme"`
		Pynb string `json:"pynb"`
		Tags string `json:"tags"`
		Category string `json:"category"`
		Image struct {
			Files []string `json:"files"`
			Cdnuri string `json:"cdnUri"`
		} `json:"image"`
		Diskspace int `json:"diskspace"`
		RAM int `json:"ram"`
		Public bool `json:"public"`
		StartCount int `json:"start_count"`
		Createdat string `json:"createdAt"`
		Parent string `json:"parent"`
		Children []string `json:"children"`
		Featured bool `json:"featured"`
		Author string `json:"author"`
		ID string `json:"id"`
		Templatekey string `json:"templateKey"`
	} `json:"snapshot"`
	Success bool `json:"success"`
}


type Output_Terminal_Access struct {
	IsPublicList []int `json:"is_public_list"`
	AccessRules []string `json:"access_rules"`
	Links []interface{} `json:"links"`
}


type Output_Cname_Records struct {
	Available []string `json:"available"`
	Assigned []struct {
		Domain string `json:"domain"`
		Subdomain string `json:"subdomain"`
		Port string `json:"port"`
	} `json:"assigned"`
}


type Output_Terminal_Idle_Settings struct {
	Success bool `json:"success"`
	Settings struct {
		Action string `json:"action"`
		Triggers struct {
		} `json:"triggers"`
	} `json:"settings"`
}


type Output_Instance_Types struct {
	InstanceTypes struct {
		Micro struct {
			CPU string `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"micro"`
		Mini struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"mini"`
		Small struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"small"`
		Medium struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"medium"`
		Xlarge struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"xlarge"`
		TwoXlarge struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"2xlarge"`
		FourXlarge struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"4xlarge"`
		EightXlarge struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"8xlarge"`
		Gpu struct {
			CPU int `json:"cpu"`
			RAM int `json:"ram"`
			Price float64 `json:"price"`
		} `json:"gpu"`
	} `json:"instance_types"`
}


type Output_Instance_Price struct {
	Price float64 `json:"price"`
	Units string `json:"units"`
}

type Output_Balance struct {
	Balance float64 `json:"balance"`
}

type Output_Balance_Added struct {
	Events []struct {
		Reason string `json:"reason"`
		Amount int `json:"amount"`
		Time int64 `json:"time"`
	} `json:"events"`
	Total int `json:"total"`
}

type Output_Burn_History struct {
	CpusHistory [][]float64 `json:"cpus_history"`
}

type Output_Burn_State struct {
	CpusState [][]float64 `json:"cpus_state"`
	Burn float64 `json:"burn"`
	Cpus float64 `json:"cpus"`
	Time int `json:"time"`
}

type Output_Burn_Estimates struct {
	Total float64 `json:"total"`
	TerminalBurnEstimates string `json:"terminal_burn_estimates"`
}

type Output_Add_Key_To_Terminal struct {
	Publickey string `json:"publicKey"`
	Success bool `json:"success"`
}

type Output_Add_Key_To_Ssh_Proxy struct {
	Publickeys []struct {
		Name string `json:"name"`
		Value string `json:"value"`
		Fingerprint string `json:"fingerprint"`
		Creationtime string `json:"creationTime"`
		ID string `json:"_id"`
	} `json:"publicKeys"`
}


type Output_Request_Progress struct {
	Operation string `json:"operation"`
	Status string `json:"status"`
	State string `json:"state"`
	Result struct {
		CPU string `json:"cpu"`
		RAM int `json:"ram"`
		Diskspace int `json:"diskspace"`
		Name string `json:"name"`
		SnapshotID string `json:"snapshot_id"`
		Temporary bool `json:"temporary"`
		Status string `json:"status"`
		AllowSpot bool `json:"allow_spot"`
		ContainerKey string `json:"container_key"`
		Subdomain string `json:"subdomain"`
		ContainerIP string `json:"container_ip"`
		CreationTime int64 `json:"creation_time"`
	} `json:"result"`
}



type Output_Snapshot_Progress_Results struct {
	Result string `json:"result"`
	Operation string `json:"operation"`
	State string `json:"state"`
	Candidate struct {
		Title string `json:"title"`
		Body string `json:"body"`
		Readme string `json:"readme"`
		Tags string `json:"tags"`
		Image struct {
			Files []string `json:"files"`
			Cdnuri string `json:"cdnUri"`
		} `json:"image"`
		Diskspace int `json:"diskspace"`
		RAM int `json:"ram"`
		Public bool `json:"public"`
		Featured bool `json:"featured"`
		Terminal struct {
			CPU string `json:"cpu"`
			RAM string `json:"ram"`
			Diskspace string `json:"diskspace"`
			Name string `json:"name"`
			SnapshotID string `json:"snapshot_id"`
			Temporary string `json:"temporary"`
			CustomData string `json:"custom_data"`
			Status string `json:"status"`
			AllowSpot string `json:"allow_spot"`
			ContainerKey string `json:"container_key"`
			Subdomain string `json:"subdomain"`
			ContainerIP string `json:"container_ip"`
			CreationTime string `json:"creation_time"`
		} `json:"terminal"`
	} `json:"candidate"`
	Status string `json:"status"`
}

type Output_Who_Am_I struct {
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		URL string `json:"url"`
		Company string `json:"company"`
		Location string `json:"location"`
		Balance float64 `json:"balance"`
		Email string `json:"email"`
		IsAdmin bool `json:"is_admin"`
		ProfileImage string `json:"profile_image"`
	} `json:"user"`
}

// DATA INPUT STRUCTS

type Input_Output_creds struct {
	Utoken string `json:"usertoken"`
	Atoken string `json:"atoken"`
}

type Input_List_Public_Snapshots struct {
	Username string `json:"username"`
	Tag	string	`json:"tag"`
	Featured bool	`json:"featured"`
	Title	string	`json:"title"`
	Page	int	`json:"page"`
	PerPage	int	`json:"perPage"`
	Sortby	string `json:"sortby"`
}

type Input_Count_Public_Snapshots struct {
	Username string `json:"username"`
	Tag	string	`json:"tag"`
	Featured bool	`json:"featured"`
	Title	string	`json:"title"`
}

type Input_Get_Terminal struct {
	ContainerKey string `json:"container_key"`
	Subdomain string `json:"subdomain"`
}

type Input_Start_Snapshot struct {
	SnapshotID	string	`json:"snapshot_id"`
	InstanceType	string	`json:"instance_type"`
	Temporary	bool	`json:"temporary"`
	Name	string	`json:"name"`
	Autopause	bool `json:"autopause"`
	StartupScript	string	`json:"startup_script"`
	KeepRam bool `json:"keep_ram"`
}

type Input_Edit_Terminal struct {
	ContainerKey string `json:"container_key"`
	InstanceType	string	`json:"instance_type"`
	Diskspace	int	`json:"diskspace"`
	Name	string	`json:"name"`
	CustomData	string `json:"custom_data"`
}

type  Input_List_Snapshots struct {
	Tag string `json:"tag"`
	Featured bool `json:"featured"`
	Title string `json:"title"`
	Page int `json:"page"`
	Perpage int `json:"perpage"`
	Sortby string `json:"sortby"`
}

type Input_Count_Snapshots struct {
	Tag	string	`json:"tag"`
	Featured bool	`json:"featured"`
	Title	string	`json:"title"`
}

type Input_Edit_Snapshot struct {
	SnapshotId string `json:"snapshot_id"`
	Body string `json:"body"`
	Title string `json:"title"`
	Readme string `json:"readme"`
	Tags string `json:"tags"`
}

type Input_Snapshot_Terminal struct {
	ContainerKey string `json:"container_key"`
	Body string `json:"body"`
	Title string `json:"title"`
	Readme string `json:"readme"`
	Tags string `json:"tags"`
	Public bool `json:"public"`
	KeepRam bool `json:"keep_ram"`
}

type Input_Edit_TerminalAccess struct {
	ContainerKey string `json:"container_key"`
	IsPublicList []int `json:"is_public_list"`
	AccessRules []string `json:"access_rules"`
}


type Input_gift struct {
	Email string	`json:"email"`
	Cents float64	`json:"cents"`
}

// USER DATA INPUT STRUCTS

type UInput_Links struct {
	Port string	`json:"port"`
	Source string `json:"source"`
}
