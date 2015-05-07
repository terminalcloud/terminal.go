package terminalgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var baseurl string = "https://api.terminal.com/v0.2/"
var utoken string
var atoken string


//// ******************** //
///      FUNCTIONS       ///
// ******************** ////


//****** API REQUEST MAKER ******//

func Make_Request(call string, kind string, data []byte) string {
	url := baseurl + call
	fmt.Println(string(call))
	fmt.Println(string(data))
	data_buffer := bytes.NewBuffer(data)
	req, _ := http.NewRequest(kind, url, data_buffer)
	req.Header.Set("user-token", utoken)
	req.Header.Set("access-token", atoken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}


//****** CREDENTIAL MANAGEMENT ******//

func Set_Credentials(file bool, filename string, u_token string, a_token string) {
	utoken = u_token
	atoken = a_token
	if file == true {
		input := &Input_Output_creds{Utoken: u_token, Atoken: a_token}
		data, _ := json.Marshal(input)
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
		output, err := f.WriteString(string(data))
		if err != nil {
			fmt.Println(output, err)
		}
		f.Close()
	}
}

func Load_Credentials(file bool, filename string) bool {
	return true
} // TODO: Make this one.


//****** BROWSE SNAPSHOTS & USERS ******//

func Get_Snapshot(snapshot_id string) *Output_Snapshot {
	call := "get_snapshot"
	input := map[string]string{"snapshot_id":snapshot_id}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Snapshot{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Get_Profile(username string) *Output_User {
	call := "get_profile"
	input := map[string]string{"username": username}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_User{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func List_Public_Snapshots(username string, tag string, featured bool, title string, page int, perPage int, sortby string) *Output_Snapshots_List {
	call := "list_public_snapshots"
	input := &Input_List_Public_Snapshots{Username: username, Tag: tag, Featured: featured, Title: title, Page: page, PerPage: perPage, Sortby: sortby}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Snapshots_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Count_Public_Snapshots(username string, tag string, featured bool, title string) *Output_Snapshot_Count {
	call := "count_public_snapshots"
	input := &Input_Count_Public_Snapshots{Username: username, Tag: tag, Featured: featured, Title: title}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Snapshot_Count{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

//****** CREATE & MANAGE TERMINALS ******//

func List_Terminals() *Terminal_List {
	output := Make_Request("list_terminals", "POST", []byte(""))
	res := &Terminal_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Get_Terminal(container_key string, subdomain string) *Output_Terminal {
	call := "get_terminal"
	input := &Input_Get_Terminal{ContainerKey: container_key, Subdomain: subdomain}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Terminal{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Start_Snapshot(snapshot_id string, instance_type string, temporary bool, name string, autopause bool, startup_script string) *Output_RequestID {
	call := "start_snapshot"
	input := &Input_Start_Snapshot{SnapshotID: snapshot_id, InstanceType: instance_type, Temporary: temporary, Name: name, Autopause: autopause, StartupScript: startup_script}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Delete_Terminal(container_key string) *Output_Req_Status {
	call := "delete_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Req_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Pause_Terminal(container_key string) *Output_RequestID {
	call := "pause_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Resume_Terminal(container_key string) *Output_RequestID {
	call := "resume_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Edit_Terminal(container_key string, instance_type string, diskspace int, name string, custom_data string) *Output_Req_Status {
	call := "edit_terminal"
	input := &Input_Edit_Terminal{ContainerKey: container_key, InstanceType: instance_type, Diskspace: diskspace, Name: name, CustomData: custom_data}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Req_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}


//****** CREATE & MANAGE SNAPSHOTS ******//

func List_Snapshots(tag string, featured bool, title string, page int, perPage int ,sortby string) *Output_Snapshots_Priv_List {
	call := "list_snapshots"
	input := &Input_List_Snapshots{Tag: tag, Featured: featured, Title: title, Page: page, Sortby: sortby}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Snapshots_Priv_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Count_Snapshots(tag string, featured bool, title string) *Output_Snapshot_Count{
	call := "count_snapshots"
	input := &Input_Count_Snapshots{Tag: tag, Featured: featured, Title: title}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	fmt.Println(string(output))
	res := &Output_Snapshot_Count{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Delete_Snapshot(snapshot_id string) *Output_Status {
	call := "delete_snapshot"
	input := map[string]string{"snapshot_id":snapshot_id}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Edit_Snapshot(snapshot_id string, body string, title string, readme string, tags string) *Output_Snapshot_Edit {
	call := "edit_snapshot"
	input := &Input_Edit_Snapshot{SnapshotId:snapshot_id, Body:body, Title:title, Readme:readme, Tags:tags}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Snapshot_Edit{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Snapshot_Terminal(container_key string, body string, title string, readme string, tags string, public bool) *Output_Req_Success {
	call := "snapshot_terminal"
	input := &Input_Snapshot_Terminal{ContainerKey:container_key, Body:body, Title:title, Readme:readme, Tags:tags}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Req_Success{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

//****** MANAGE TERMINAL ACCESS ******//

func Add_Terminal_Links(container_key string, links []UInput_Links) *Output_Status {
	call:= "add_terminal_links"
	slinks, _ := json.Marshal(links)
	input:= fmt.Sprintf("{{\"container_id\":\"%s\"}, \"links\": %s}", container_key, slinks)
	data := []byte(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Remove_Terminal_Links(container_key string, links []UInput_Links) *Output_Status {
	call := "remove_terminal_links"
	slinks, _ := json.Marshal(links)
	input := fmt.Sprintf("{{\"container_id\":\"%s\"}, \"links\": %s}", container_key, slinks)
	data := []byte(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func List_Terminal_Access(container_key string) *Output_Terminal_Access{
	call := "list_terminal_access"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Terminal_Access{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Edit_Terminal_Access(container_key string, is_public_list []int, access_rules []string) *Output_Status {
	call := "edit_terminal_access"
	input := &Input_Edit_TerminalAccess{ContainerKey: container_key, IsPublicList: is_public_list, AccessRules:access_rules}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}


//****** MANAGE TERMINAL DNS & DOMAINS ******//

func Get_Cname_Records() *Output_Cname_Records {
	call := "get_cname_records"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Add_Domain_To_Pool(domain string) *Output_Cname_Records {
	call := "add_domain_to_pool"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Remove_Domain_From_Pool(domain string) *Output_Cname_Records {
	call := "remove_domain_from_pool"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Add_Cname_Record(cname string, subdomain string, port string) *Output_Cname_Records {
	call := "add_cname_record"
	input := map[string]string{"domain": cname, "subdomain": subdomain, "port": port}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Remove_Cname_Record(domain string) *Output_Cname_Records {
	call := "remove_cname_record"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}


//****** MANAGE TERMINAL IDLE SETTINGS ******//

func Set_Terminal_Idle_Settings(container_key string, triggers string, action string) *Output_Terminal_Iddle_Settings {
	call := "set_terminal_idle_settings"
	input := map[string]string{"container_key": container_key, "triggers": triggers, "action": action}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Terminal_Iddle_Settings{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Get_Terminal_Idle_Settings(container_key string) *Output_Terminal_Iddle_Settings {
	call := "get_terminal_idle_settings"
	input := map[string]string{"container_key": container_key}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Terminal_Iddle_Settings{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

//****** MANAGE USAGE AND CREDITS ******//

func Instance_Types() *Output_Instance_Types {
	call := "instance_types"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Instance_Types{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Instance_Price(instance_type string) *Output_Instance_Price {
	call := "instance_Price"
	input := map[string]string{"instance_type": instance_type}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Instance_Price{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Balance() *Output_Balance {
	call := "balance"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Balance{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Balance_Added() *Output_Balance_Added {
	call := "balance_added"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Balance_Added{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Gift(email string, cents float64) *Output_Status {
	call := "gift"
	input := &Input_gift{Email: email, Cents: cents}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	fmt.Println(string(output))
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Burn_History() *Output_Burn_History {
	call := "burn_history"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Burn_History{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Terminal_Usage_History() string { // TODO: Check this Function
	call := "terminal_usage_history"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := string(output)
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Burn_State() *Output_Burn_State {
	call := "burn_state"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Burn_State{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Burn_Estimates() *Output_Burn_Estimates { // TODO: Check this Function
	call := "burn_estimates"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Burn_Estimates{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}


//****** MANAGE SSH PUBLIC KEYS ******//

func Add_Authorized_Key_To_Terminal(container_key string, publicKey string) *Output_Add_Key_To_Terminal {
	call := "add_authorized_key_to_terminal"
	input := map[string]string{"container_key": container_key, "publicKey": publicKey}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Add_Key_To_Terminal{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Add_Authorized_Key_To_Ssh_Proxy(name string, publicKey string) *Output_Add_Key_To_Ssh_Proxy {
	call := "add_authorized_key_to_ssh_proxy"
	input := map[string]string{"name": name, "publicKey": publicKey}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}


func Del_Authorized_Key_From_Ssh_Proxy(name string, fingerprint string) *Output_Add_Key_To_Ssh_Proxy {
	call := "del_authorized_key_from_ssh_proxy"
	input := map[string]string{"name": name,"fingerprint": fingerprint}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Get_Authorized_Key_From_Ssh_Proxy() *Output_Add_Key_To_Ssh_Proxy {
	call := "get_authorized_keys_from_ssh_proxy"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

//****** OTHER ******//

func Request_Progress(request_id string) *Output_Request_Progress {
	call := "request_progress"
	input := map[string]string{"request_id": request_id}
	data, _ := json.Marshal(input)
	output := Make_Request(call, "POST", data)
	res := &Output_Request_Progress{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}

func Who_Am_I() *Output_Who_Am_I {
	call := "who_am_i"
	data := []byte("")
	output := Make_Request(call, "POST", data)
	res := &Output_Who_Am_I{}
	json.Unmarshal([]byte(string(output)), &res)
	return res
}
