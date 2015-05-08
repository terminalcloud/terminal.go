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

func Make_Request(call string, kind string, data []byte) (string, int) {
	url := baseurl + call
	fmt.Println(string(call))
	fmt.Println(string(data))
	data_buffer := bytes.NewBuffer(data)
	req, err := http.NewRequest(kind, url, data_buffer)
	if err != nil {panic(err)}
	req.Header.Set("user-token", utoken)
	req.Header.Set("access-token", atoken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {panic(err)}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {panic(err)}
	return string(body), response.StatusCode
}


//****** CREDENTIAL MANAGEMENT ******//

func Set_Credentials(file bool, filename string, u_token string, a_token string) {
	utoken = u_token
	atoken = a_token
	if file == true {
		input := &Input_Output_creds{Utoken: u_token, Atoken: a_token}
		data, _ := json.Marshal(input)
		f, err := os.Create(filename)
		if err != nil {panic(err)}
		_ , err = f.WriteString(string(data))
		if err != nil {panic(err)}
		f.Close()
	}
}

func Load_Credentials() string {
	_, err := os.Create("/caca")
	if err != nil {panic(err)}
	return "returned"

} // TODO: Make this one.


//****** BROWSE SNAPSHOTS & USERS ******//

func Get_Snapshot(snapshot_id string) (*Output_Snapshot, bool) {
	output, success  := Get_Snapshot_Raw(snapshot_id)
	res := &Output_Snapshot{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Snapshot_Raw(snapshot_id string) (string, bool) {
	call := "get_snapshot"
	input := map[string]string{"snapshot_id":snapshot_id}
	data, _ := json.Marshal(input)
	output,  status_code  := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Get_Profile(username string) (*Output_User, bool) {
	output, success  := Get_Profile_Raw(username)
	res := &Output_User{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Profile_Raw(username string) (string, bool) {
	call := "get_profile"
	input := map[string]string{"username": username}
	data, _ := json.Marshal(input)
	output, status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func List_Public_Snapshots(username string, tag string, featured bool, title string, page int, perPage int, sortby string) (*Output_Snapshots_List, bool) {
	output, success  := List_Public_Snapshots_Raw(username, tag, featured, title, page, perPage, sortby)
	res := &Output_Snapshots_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func List_Public_Snapshots_Raw(username string, tag string, featured bool, title string, page int, perPage int, sortby string) (string, bool) {
	call := "list_public_snapshots"
	input := &Input_List_Public_Snapshots{Username: username, Tag: tag, Featured: featured, Title: title, Page: page, PerPage: perPage, Sortby: sortby}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Count_Public_Snapshots(username string, tag string, featured bool, title string) (*Output_Snapshot_Count, bool) {
	output, success  := Count_Public_Snapshots_Raw(username, tag, featured, title)
	res := &Output_Snapshot_Count{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Count_Public_Snapshots_Raw(username string, tag string, featured bool, title string) (string, bool) {
	call := "count_public_snapshots"
	input := &Input_Count_Public_Snapshots{Username: username, Tag: tag, Featured: featured, Title: title}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** CREATE & MANAGE TERMINALS ******//

func List_Terminals() (*Terminal_List, bool) {
	output, success  := List_Terminals_Raw()
	res := &Terminal_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func List_Terminals_Raw() (string, bool) {
	output, status_code := Make_Request("list_terminals", "POST", []byte(""))
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Get_Terminal(container_key string, subdomain string) (*Output_Terminal, bool) {
	output, success  := Get_Terminal_Raw(container_key, subdomain)
	res := &Output_Terminal{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Terminal_Raw(container_key string, subdomain string) (string, bool) {
	call := "get_terminal"
	input := &Input_Get_Terminal{ContainerKey: container_key, Subdomain: subdomain}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Start_Snapshot(snapshot_id string, instance_type string, temporary bool, name string, autopause bool, startup_script string) (*Output_RequestID, bool) {
	output, success  := Start_Snapshot_Raw(snapshot_id, instance_type, temporary, name, autopause, startup_script)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Start_Snapshot_Raw(snapshot_id string, instance_type string, temporary bool, name string, autopause bool, startup_script string) (string, bool) {
	call := "start_snapshot"
	input := &Input_Start_Snapshot{SnapshotID: snapshot_id, InstanceType: instance_type, Temporary: temporary, Name: name, Autopause: autopause, StartupScript: startup_script}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Delete_Terminal(container_key string) (*Output_Req_Status, bool) {
	output, success  := Delete_Terminal_Raw(container_key)
	res := &Output_Req_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Delete_Terminal_Raw(container_key string) (string, bool) {
	call := "delete_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Pause_Terminal(container_key string) (*Output_RequestID, bool) {
	output, success  := Pause_Terminal_Raw(container_key)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Pause_Terminal_Raw(container_key string) (string, bool) {
	call := "pause_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Resume_Terminal(container_key string) (*Output_RequestID, bool) {
	output, success  := Resume_Terminal_Raw(container_key)
	res := &Output_RequestID{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Resume_Terminal_Raw(container_key string) (string, bool) {
	call := "resume_terminal"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Edit_Terminal(container_key string, instance_type string, diskspace int, name string, custom_data string) (*Output_Req_Status, bool) {
	output, success  := Edit_Terminal_Raw(container_key, instance_type, diskspace, name, custom_data)
	res := &Output_Req_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Edit_Terminal_Raw(container_key string, instance_type string, diskspace int, name string, custom_data string) (string, bool) {
	call := "edit_terminal"
	input := &Input_Edit_Terminal{ContainerKey: container_key, InstanceType: instance_type, Diskspace: diskspace, Name: name, CustomData: custom_data}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


//****** CREATE & MANAGE SNAPSHOTS ******//

func List_Snapshots(tag string, featured bool, title string, page int, perPage int ,sortby string) (*Output_Snapshots_Priv_List, bool) {
	output, success  := List_Snapshots_Raw(tag, featured, title, page, perPage ,sortby)
	res := &Output_Snapshots_Priv_List{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func List_Snapshots_Raw(tag string, featured bool, title string, page int, perPage int ,sortby string) (string, bool) {
	call := "list_snapshots"
	input := &Input_List_Snapshots{Tag: tag, Featured: featured, Title: title, Page: page, Sortby: sortby}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}



func Count_Snapshots(tag string, featured bool, title string) (*Output_Snapshot_Count, bool){
	output, success  := Count_Snapshots_Raw(tag, featured, title)
	res := &Output_Snapshot_Count{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Count_Snapshots_Raw(tag string, featured bool, title string) (string, bool){
	call := "count_snapshots"
	input := &Input_Count_Snapshots{Tag: tag, Featured: featured, Title: title}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Delete_Snapshot(snapshot_id string) (*Output_Status, bool){
	output, success  := Delete_Snapshot_Raw(snapshot_id)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}


func Delete_Snapshot_Raw(snapshot_id string) (string, bool) {
	call := "delete_snapshot"
	input := map[string]string{"snapshot_id":snapshot_id}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Edit_Snapshot(snapshot_id string, body string, title string, readme string, tags string) (*Output_Snapshot_Edit, bool) {
	output, success  := Edit_Snapshot_Raw(snapshot_id, body, title, readme, tags)
	res := &Output_Snapshot_Edit{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Edit_Snapshot_Raw(snapshot_id string, body string, title string, readme string, tags string) (string, bool) {
	call := "edit_snapshot"
	input := &Input_Edit_Snapshot{SnapshotId:snapshot_id, Body:body, Title:title, Readme:readme, Tags:tags}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Snapshot_Terminal(container_key string, body string, title string, readme string, tags string, public bool) (*Output_Req_Success, bool) {
	output, success  := Snapshot_Terminal_Raw(container_key, body, title, readme, tags, public)
	res := &Output_Req_Success{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Snapshot_Terminal_Raw(container_key string, body string, title string, readme string, tags string, public bool) (string, bool) {
	call := "snapshot_terminal"
	input := &Input_Snapshot_Terminal{ContainerKey:container_key, Body:body, Title:title, Readme:readme, Tags:tags}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** MANAGE TERMINAL ACCESS ******//

func Add_Terminal_Links(container_key string, links []UInput_Links) (*Output_Status, bool) {
	output, success  := Add_Terminal_Links_Raw(container_key, links)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Add_Terminal_Links_Raw(container_key string, links []UInput_Links) (string, bool) {
	call:= "add_terminal_links"
	slinks, _ := json.Marshal(links)
	input:= fmt.Sprintf("{{\"container_id\":\"%s\"}, \"links\": %s}", container_key, slinks)
	data := []byte(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Remove_Terminal_Links(container_key string, links []UInput_Links) (*Output_Status, bool) {
	output, success  := Remove_Terminal_Links_Raw(container_key, links)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Remove_Terminal_Links_Raw(container_key string, links []UInput_Links) (string, bool) {
	call := "remove_terminal_links"
	slinks, _ := json.Marshal(links)
	input := fmt.Sprintf("{{\"container_id\":\"%s\"}, \"links\": %s}", container_key, slinks)
	data := []byte(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func List_Terminal_Access(container_key string) (*Output_Terminal_Access, bool) {
	output, success  := List_Terminal_Access_Raw(container_key)
	res := &Output_Terminal_Access{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func List_Terminal_Access_Raw(container_key string) (string, bool) {
	call := "list_terminal_access"
	input := map[string]string{"container_key":container_key}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Edit_Terminal_Access(container_key string, is_public_list []int, access_rules []string) (*Output_Status, bool) {
	output, success := Edit_Terminal_Access_Raw(container_key, is_public_list, access_rules)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Edit_Terminal_Access_Raw(container_key string, is_public_list []int, access_rules []string) (*Output_Status, bool) {
	call := "edit_terminal_access"
	input := &Input_Edit_TerminalAccess{ContainerKey: container_key, IsPublicList: is_public_list, AccessRules:access_rules}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** MANAGE TERMINAL DNS & DOMAINS ******//

func Get_Cname_Records() (*Output_Cname_Records, bool) {
	output, success  := Get_Cname_Records_Raw()
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Cname_Records_Raw() (string, bool) {
	call := "get_cname_records"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Add_Domain_To_Pool(domain string) (*Output_Cname_Records, bool) {
	output, success := Add_Domain_To_Pool_Raw(domain)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Add_Domain_To_Pool_Raw(domain string) (string, bool) {
	call := "add_domain_to_pool"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Remove_Domain_From_Pool(domain string) (*Output_Cname_Records, bool) {
	output, success  := Remove_Domain_From_Pool_Raw(domain)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Remove_Domain_From_Pool_Raw(domain string) (string, bool) {
	call := "remove_domain_from_pool"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Add_Cname_Record(cname string, subdomain string, port string) (*Output_Cname_Records, bool) {
	output, success := Add_Cname_Record_Raw(cname, subdomain, port)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Add_Cname_Record_Raw(cname string, subdomain string, port string) (string, bool) {
	call := "add_cname_record"
	input := map[string]string{"domain": cname, "subdomain": subdomain, "port": port}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Remove_Cname_Record(domain string) (*Output_Cname_Records, bool) {
	output, success := Remove_Cname_Record_Raw(domain)
	res := &Output_Cname_Records{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Remove_Cname_Record_Raw(domain string) (string, bool) {
	call := "remove_cname_record"
	input := map[string]string{"domain": domain}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** MANAGE TERMINAL IDLE SETTINGS ******//

func Set_Terminal_Idle_Settings(container_key string, triggers string, action string) (*Output_Terminal_Idle_Settings, bool) {
	output, success := Set_Terminal_Idle_Settings_Raw(container_key, triggers, action)
	res := &Output_Terminal_Idle_Settings{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Set_Terminal_Idle_Settings_Raw(container_key string, triggers string, action string) (string, bool) {
	call := "set_terminal_idle_settings"
	input := map[string]string{"container_key": container_key, "triggers": triggers, "action": action}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Get_Terminal_Idle_Settings(container_key string) (*Output_Terminal_Idle_Settings, bool) {
	output, success := Get_Terminal_Idle_Settings_Raw(container_key)
	res := &Output_Terminal_Idle_Settings{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Terminal_Idle_Settings_Raw(container_key string) (string, bool) {
	call := "get_terminal_idle_settings"
	input := map[string]string{"container_key": container_key}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** MANAGE USAGE AND CREDITS ******//

func Instance_Types() (*Output_Instance_Types, bool) {
	output, success := Instance_Types_Raw()
	res := &Output_Instance_Types{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Instance_Types_Raw() (string, bool) {
	call := "instance_types"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Instance_Price(instance_type string) (*Output_Instance_Price, bool) {
	output, success := Instance_Price_Raw(instance_type)
	res := &Output_Instance_Price{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Instance_Price_Raw(instance_type string) (string, bool) {
	call := "instance_Price"
	input := map[string]string{"instance_type": instance_type}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Balance() (*Output_Balance, bool) {
	output, success := Balance_Raw()
	res := &Output_Balance{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Balance_Raw() (string, bool) {
	call := "balance"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Balance_Added() (*Output_Balance_Added, bool) {
	output, success := Balance_Added_Raw()
	res := &Output_Balance_Added{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Balance_Added_Raw() (string, bool) {
	call := "balance_added"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Gift(email string, cents float64) (*Output_Status, bool) {
	output, success := Gift_Raw(email, cents)
	res := &Output_Status{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Gift_Raw(email string, cents float64) (string, bool) {
	call := "gift"
	input := &Input_gift{Email: email, Cents: cents}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Burn_History() (*Output_Burn_History, bool) {
	output, success := Burn_History_Raw()
	res := &Output_Burn_History{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Burn_History_Raw() (string, bool) {
	call := "burn_history"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Terminal_Usage_History_Raw() (string, bool) { // TODO: Non-Raw version
	call := "terminal_usage_history"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Burn_State() (*Output_Burn_State, bool) {
	output, success := Burn_State_Raw()
	res := &Output_Burn_State{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}


func Burn_State_Raw() (string, bool) {
	call := "burn_state"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Burn_Estimates() (*Output_Burn_Estimates, bool) {
	output, success :=  Burn_Estimates_Raw()
	res := &Output_Burn_Estimates{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Burn_Estimates_Raw() (string, bool) {
	call := "burn_estimates"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** MANAGE SSH PUBLIC KEYS ******//

func Add_Authorized_Key_To_Terminal(container_key string, publicKey string) (*Output_Add_Key_To_Terminal, bool) {
	output, success := Add_Authorized_Key_To_Terminal_Raw(container_key, publicKey)
	res := &Output_Add_Key_To_Terminal{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Add_Authorized_Key_To_Terminal_Raw(container_key string, publicKey string) (string, bool) {
	call := "add_authorized_key_to_terminal"
	input := map[string]string{"container_key": container_key, "publicKey": publicKey}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Add_Authorized_Key_To_Ssh_Proxy(name string, publicKey string) (*Output_Add_Key_To_Ssh_Proxy, bool) {
	output, success := Add_Authorized_Key_To_Ssh_Proxy_Raw(name, publicKey)
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Add_Authorized_Key_To_Ssh_Proxy_Raw(name string, publicKey string) (string, bool) {
	call := "add_authorized_key_to_ssh_proxy"
	input := map[string]string{"name": name, "publicKey": publicKey}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Del_Authorized_Key_From_Ssh_Proxy(name string, fingerprint string) (*Output_Add_Key_To_Ssh_Proxy, bool) {
	output, success := Del_Authorized_Key_From_Ssh_Proxy_Raw(name, fingerprint)
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Del_Authorized_Key_From_Ssh_Proxy_Raw(name string, fingerprint string) (string, bool) {
	call := "del_authorized_key_from_ssh_proxy"
	input := map[string]string{"name": name,"fingerprint": fingerprint}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

func Get_Authorized_Key_From_Ssh_Proxy() (*Output_Add_Key_To_Ssh_Proxy, bool) {
	output, success := Get_Authorized_Key_From_Ssh_Proxy_Raw()
	res := &Output_Add_Key_To_Ssh_Proxy{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Get_Authorized_Key_From_Ssh_Proxy_Raw() (string, bool) {
	call := "get_authorized_keys_from_ssh_proxy"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}

//****** OTHER ******//

func Request_Progress(request_id string) (*Output_Request_Progress, bool) {
	output, success := Request_Progress_Raw(request_id)
	res := &Output_Request_Progress{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Request_Progress_Raw(request_id string) (string, bool) {
	call := "request_progress"
	input := map[string]string{"request_id": request_id}
	data, _ := json.Marshal(input)
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}


func Who_Am_I() (*Output_Who_Am_I, bool) {
	output, success := Who_Am_I_Raw()
	res := &Output_Who_Am_I{}
	json.Unmarshal([]byte(string(output)), &res)
	return res, success
}

func Who_Am_I_Raw() (string, bool) {
	call := "who_am_i"
	data := []byte("")
	output,  status_code := Make_Request(call, "POST", data)
	success := false
	if status_code == 200 { success = true }
	res := string(output)
	return res, success
}
