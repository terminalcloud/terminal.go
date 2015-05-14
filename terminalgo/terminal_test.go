package terminalgo

import (
  	"testing"
	"time"
	"fmt"
)

// Test Configuration
var creds_file string = "creds.json"
var username string = "terminal"
var test_name string = "golang_tests"
var snapshot_id string = "987f8d702dc0a6e8158b48ccd3dec24f819a7ccb2756c396ef1fd7f5b34b7980" // Ubuntu Snap
var err error
var testdomain string = "golangtestdomain.com"
var publicKey string = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDbM1JpfSiRtFwNuKVyqA8UuaFedjDf/eYLoWm4OwcpdQw4070FLx4yDTNj6GFyaB+cOMFzTMS5eWEpnqKlcDe52IX74hR+XXP6/Cd6PhXJ8ABe30YCxEsKe48VgxRt5c+zYwjwi7IQmXNkOOf2MHXIBD7BnyWVhM9wu/W4PegCl0b4p95G78L3Kv+uUJOE1Y0Ue4r7mn9ZoiI16EBP9yk2ptVAVCwVurCALEFIJpZ3lG8fLLqEpiUdOsHym/ugW9CuvneiywQpRoaz3oH1NGBJHKpFyC7KjYYzVkSYMtMc4WWSMOKwqP+J5TXDhb+KCDYrFGMYgCuO+APGzck2KGKL"
var publicKey_fingerprint string = "f8:cb:db:e1:ab:16:97:eb:57:29:15:82:e6:4e:35:3b"

func Test_Load_Credentials(t *testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	if err != nil {t.Error(err)}
}

func Test_Write_Credentials(t *testing.T) {
	err = Write_Credentials(creds_file, utoken, atoken)
	if err != nil {t.Error(err)}
}

func Test_Who_Am_I(t *testing.T) {
	res, success := Who_Am_I_Raw()
	if success != true {t.Error(res)}
}


func Test_Get_Snapshot(t *testing.T) {
	res, success := Get_Snapshot_Raw(snapshot_id)
	if success != true {t.Error(res)}
}

func Test_Get_Profile(t *testing.T) {
	res, success := Get_Profile_Raw(username)
	if success != true {t.Error(res)}
}

func Test_List_Terminals(t *testing.T) {
	res, success := List_Terminals_Raw()
	if success != true {t.Error(res)}
}


func Test_Public_Snapshots_Listing(t *testing.T) {
	res_raw, success := List_Public_Snapshots_Raw(username,"ubuntu",true,"",0,100,"popularity")
	if success != true {t.Error("--- FAIL: List_Public_Snapshots", res_raw)} else {fmt.Println("--- PASS: List_Public_Snapshots")}
	res_raw, success = Count_Public_Snapshots_Raw(username, "ubuntu", true, "")
	if success != true {t.Error("--- FAIL: Count_Public_Snapshots", res_raw)} else {fmt.Println("--- PASS: Count_Public_Snapshots")}
}


// Terminal Handling Tests

func Test_General_Terminal_Handling(t *testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res, success := Start_Snapshot(snapshot_id, "micro", false, test_name, true, "")
	if success != true {t.Error("Start_Snapshot failing, Aborting current test")}
	fmt.Println("--- PASS: Start_Snapshot")
	req_id := string(res.RequestID)
	prog, success_prog := Request_Progress(req_id)
	fmt.Println("--- PASS: Request_Progress")
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	time.Sleep( 2 * time.Second)
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	if success_prog != true { t.Error("--- FAIL: Effective snapshot start [TIMEOUT]")} else {fmt.Println("--- PASS: Effective snapshot start")}
	res_getterm, success := Get_Terminal_Raw(ckey,"")
	if success != true {t.Error(res_getterm)}
	res_edit, success := Edit_Terminal_Raw(ckey, "small", 15, "golang_test_edited","")
	if success != true {t.Error("-- FAIL: Cannot Edit Terminal",res_edit)} else {fmt.Println("--- PASS: Edit_Terminal")}
	time.Sleep(3 * time.Second)
	res_pause, success := Pause_Terminal(ckey)
	if success != true {t.Error("--- FAIL: Pause_Terminal ")} else {fmt.Println("--- PASS: Pause_Terminal")}
	req_id = res_pause.RequestID
	prog, success_prog = Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	res_resume, success := Resume_Terminal_Raw(ckey)
	if success != true {t.Error("--- FAIL: Resume_Terminal ",res_resume)} else {fmt.Println("--- PASS: Resume_Terminal")}
	time.Sleep(1500 * time.Millisecond)
	res_delete, success := Delete_Terminal_Raw(ckey)
	if success != true {t.Error("--- FAIL: Delete_Terminal ",res_delete)} else {fmt.Println("--- PASS: Delete_Terminal")}
}


// Snapshot Handling Tests

func Test_List_Snapshots(t * testing.T) {
	res, success := List_Snapshots_Raw("ubuntu", true, "", 0, 20, "popularity")
	if success != true {t.Error(res)}
}

func Test_Count_Snapshots(t * testing.T){
	res, success := Count_Public_Snapshots_Raw("terminal","ubuntu",true, "")
	if success != true {t.Error(res)}
}

func Test_Snapshot_Handling(t * testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res, success := Start_Snapshot(snapshot_id, "micro", true, test_name, true, "")
	if success != true {t.Error("Start_Snapshot failing, Aborting current test")}
	req_id := res.RequestID
	time.Sleep(2 * time.Second)
	prog, success_prog := Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	time.Sleep(2 * time.Second)
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	res_snap, success := Snapshot_Terminal(ckey,"body","golang_test_snap","readme","test",false )
	if success != true {t.Error("--- FAIL: Snapshot_Terminal ",res_snap)} else {fmt.Println("--- PASS: Snapshot_Terminal")}
	time.Sleep(3 * time.Second)
	req_id = res_snap.RequestID
	prog, success_prog = Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	prog_snap, success_prog := Request_Progress_Snapshot(req_id)
	if success_prog != true {t.Error("--- FAIL: Snapshot_Request_Progress")} else {fmt.Println("--- PASS: Snapshot_Request_Progress")}
	snap_id := prog_snap.Result
	res_raw, success := Edit_Snapshot_Raw(snap_id,"Changed Body", "golang_test_snap_changed","readme string","changed, ubuntu")
	if success != true {t.Error("--- FAIL: Edit_Snapshot",res_raw)} else {fmt.Println("--- PASS: Edit_Snapshot")}
	time.Sleep(2 * time.Second)
	res_raw, success = Delete_Snapshot_Raw(snap_id)
	if success != true {t.Error("--- FAIL: Delete_Snapshot",res_raw)} else {fmt.Println("--- PASS: Delete_Snapshot")}
	res_raw, success = Delete_Terminal_Raw(ckey)
	if success != true {t.Error("--- FAIL: Delete_Terminal - Cannot delete test terminal",res_raw)}
}

func Test_Terminal_Access_and_Linking(t * testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res, success := Start_Snapshot(snapshot_id, "micro", false, test_name, true, "")
	if success != true {t.Error("Start_Snapshot failing, Aborting current test")}
	req_id := res.RequestID
	time.Sleep(2 * time.Second)
	prog, success_prog := Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	time.Sleep(2 * time.Second)
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	links := []UInput_Links{UInput_Links{"8000", "terminal719"}}
	res_raw, success := Add_Terminal_Links_Raw(ckey, links)
	if success != true {t.Error("--- FAIL: Add_Terminal_Links", res_raw)} else {fmt.Println("--- PASS: Add_Terminal_Links")}
	time.Sleep(2 * time.Second)
	res_raw, success = Remove_Terminal_Links_Raw(ckey, links)
	if success != true {t.Error("--- FAIL: Remove_Terminal_Links", res_raw)} else {fmt.Println("--- PASS: Remove_Terminal_Links")}
	res_raw, success = List_Terminal_Access_Raw(ckey)
	if success != true {t.Error("--- FAIL: List_Terminal_Access", res_raw)} else {fmt.Println("--- PASS: List_Terminal_Access")}
	Delete_Terminal_Raw(ckey)
}

func Test_Domain_Management(t * testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res_raw, success := Get_Cname_Records_Raw()
	if success != true {t.Error("--- FAIL: Get_Cname_Records", res_raw)} else {fmt.Println("--- PASS: Get_Cname_Records")}
	res_raw, success = Add_Domain_To_Pool_Raw(testdomain)
	if success != true {t.Error("--- FAIL:  Add_Domain_To_Pool_Raw", res_raw)} else {fmt.Println("--- PASS: Add_Domain_To_Pool_Raw")}
	res, success := Start_Snapshot(snapshot_id, "micro", false, test_name, true, "")
	if success != true {t.Error("Start_Snapshot failing, Aborting current test")}
	req_id := res.RequestID
	time.Sleep(2 * time.Second)
	prog, success_prog := Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	time.Sleep(2 * time.Second)
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	subdomain := prog.Result.Subdomain
	res_cname, success := Add_Cname_Record(testdomain, subdomain, "80")
	if success != true {t.Error("--- FAIL:  Add_Cname_Record", res_cname)} else {fmt.Println("--- PASS: Add_Cname_Record")}
	res_raw, success = Remove_Cname_Record_Raw(testdomain)
	if success != true {t.Error("--- FAIL: Remove_Cname_Record", res_cname)} else {fmt.Println("--- PASS: Remove_Cname_Record")}
	res_raw, success = Remove_Domain_From_Pool_Raw("golangtestdomain.com")
	if success != true {t.Error("--- FAIL:  Remove_Domain_From_Pool_Raw", res_raw)} else {fmt.Println("--- PASS: Remove_Domain_From_Pool_Raw")}
	Delete_Terminal_Raw(ckey)
}


func Test_Terminal_Idle_Settings_Management(t * testing.T){
	utoken, atoken, err = Load_Credentials("creds.json")
	res, success := Start_Snapshot(snapshot_id, "micro", false, test_name, true, "")
	if success != true {t.Error("Start_Snapshot failing, Aborting current test")}
	req_id := res.RequestID
	time.Sleep(2 * time.Second)
	prog, success_prog := Request_Progress(req_id)
	if success_prog != true {t.Error(res)}
	for i := 0 ; i < 30 ; i++ {
		time.Sleep(2 * time.Second)
		prog, success_prog := Request_Progress(req_id)
		if success_prog != true {t.Error(res)}
		if prog.Status == "success" {break}
	}
	time.Sleep(2 * time.Second)
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	res_raw, success := Get_Terminal_Idle_Settings_Raw(ckey)
	if success != true {t.Error("--- FAIL: Get_Terminal_Idle_Settings", res_raw)} else {fmt.Println("--- PASS: Get_Terminal_Idle_Settings")}
	res_raw, success = Set_Terminal_Idle_Settings_Raw(ckey, "trigger", "pause")
	if success != true {t.Error("--- FAIL: Set_Terminal_Idle_Settings", res_raw)} else {fmt.Println("--- PASS: Set_Terminal_Idle_Settings")}
	Delete_Terminal_Raw(ckey)
}

func Test_Usage_And_Credits_Management(t * testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res_raw, success := Instance_Types_Raw()
	if success != true {t.Error("--- FAIL: Instance_Types", res_raw)} else {fmt.Println("--- PASS: Instance_Types")}
	res_raw, success = Instance_Price_Raw("micro")
	if success != true {t.Error("--- FAIL: Instance_Price", res_raw)} else {fmt.Println("--- PASS: Instance_Price")}
	res_raw, success = Balance_Raw()
	if success != true {t.Error("--- FAIL: Balance", res_raw)} else {fmt.Println("--- PASS: Balance")}
	res_raw, success = Balance_Added_Raw()
	if success != true {t.Error("--- FAIL: Balance_Added", res_raw)} else {fmt.Println("--- PASS: Balance_Added")}
	res_raw, success = Burn_History_Raw()
	if success != true {t.Error("--- FAIL: Burn_History", res_raw)} else {fmt.Println("--- PASS: Burn_History")}
	res_raw, success = Burn_Estimates_Raw()
	if success != true {t.Error("--- FAIL: Burn_Estimates", res_raw)} else {fmt.Println("--- PASS: Burn_Estimates")}
}

func Test_Ssh_Keys_Management(t * testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	res_raw, success := Add_Authorized_Key_To_Ssh_Proxy_Raw("golang_test", publicKey)
	if success != true {t.Error("--- FAIL: Add_Authorized_Key_To_Ssh_Proxy", res_raw)} else {fmt.Println("--- PASS: Add_Authorized_Key_To_Ssh_Proxy")}
	res_raw, success = Get_Authorized_Key_From_Ssh_Proxy_Raw()
	if success != true {t.Error("--- FAIL: Get_Authorized_Key_From_Ssh_Proxy", res_raw)} else {fmt.Println("--- PASS: Get_Authorized_Key_From_Ssh_Proxy")}
	res_raw, success = Del_Authorized_Key_From_Ssh_Proxy_Raw("golang_test",publicKey_fingerprint)
	if success != true {t.Error("--- FAIL: Del_Authorized_Key_From_Ssh_Proxy", res_raw)} else {fmt.Println("--- PASS: Del_Authorized_Key_From_Ssh_Proxy")}
}
