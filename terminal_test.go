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

func Test_Load_Credentials(t *testing.T) {
	utoken, atoken, err = Load_Credentials("creds.json")
	if err != nil {t.Error(err)}
}

func Test_Write_Credentials(t *testing.T) {
	err = Write_Credentials(creds_file, utoken, atoken)
	if err != nil {t.Error(err)}
}

func Test_Get_Snapshot(t *testing.T) {
	res, success := Get_Snapshot_Raw(snapshot_id)
	if success != true {
		t.Error(res)
	}
}


func Test_Get_Profile(t *testing.T) {
	res, success := Get_Profile_Raw(username)
	if success != true {
		t.Error(res)
	}
}

func Test_List_Public_Snapshots(t *testing.T) {
	res, success := List_Public_Snapshots_Raw(username,"ubuntu",true,"",0,100,"popularity")
	if success != true {
		t.Error(res)
	}
}

func Test_Count_Public_Snapshots(t *testing.T) {
	res, success := Count_Public_Snapshots_Raw(username, "ubuntu", true, "")
	if success != true {
		t.Error(res)
	}
}

func Test_List_Terminals(t *testing.T) {
	res, success := List_Terminals_Raw()
	if success != true {
		t.Error(res)
	}
}


// Terminal Handling Tests

func Test_General_Terminal_Handling(t *testing.T) {
	res, success := Start_Snapshot(snapshot_id, "micro", true, test_name, true, "")
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
	prog, success_prog = Request_Progress(req_id)
	ckey := prog.Result.ContainerKey
	if success_prog != true { t.Error("--- FAIL: Effective snapshot start [TIMEOUT]")} else {fmt.Println("--- PASS: Effective snapshot start")}
	res_getterm, success := Get_Terminal_Raw(ckey,"")
	if success != true {t.Error(res_getterm)}
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
}
