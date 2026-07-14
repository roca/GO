package scan_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/cobra/pScan/scan"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}

	if ps.Open.String() != "closed" {
		t.Errorf("Expected %q, got %q instead\n", "closed", ps.Open.String())
	}

	ps.Open = true
	if ps.Open.String() != "open" {
		t.Errorf("Expected %q, got %q instead\n", "open", ps.Open.String())
	}
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name          string
		expectedState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}
	host := "localhost"
	hl := &scan.HostsList{}
	hl.Add(host)
	ports := []int{}

	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}
		port, err := strconv.Atoi(portStr)
		//t.Log(portStr)
		if err != nil {
			t.Fatal(err)
		}

		ports = append(ports, port)
		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}

	res := scan.Run(hl, ports)

	if len(res) != 1 {
		t.Fatalf("Expected 1 results, got %d instead\n", len(res))
	}

	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q instead\n", host, res[0].Host)
	}

	if res[0].NotFound {
		t.Errorf("Expected host %q to be found, got NotFound instead\n", host)
	}

	if len(res[0].PortStates) != len(ports) {
		t.Fatalf("Expected %d port states, got %d instead\n", len(ports), len(res[0].PortStates))
	}

	for i, tc :=  range testCases {
		if res[0].PortStates[i].Port != ports[i] {
			t.Errorf("Expected port %d, got %d instead\n", ports[i], res[0].PortStates[i].Port)
		}

		if res[0].PortStates[i].Open.String() != tc.expectedState {
			t.Errorf("Expected port %d to be %q, got %q instead\n", ports[i], tc.expectedState, res[0].PortStates[i].Open.String())
		}
	}

}

func TestRunHostNotFound(t *testing.T) {
	host := "389.389.389.389"
	hl := &scan.HostsList{}
	hl.Add(host)

	res := scan.Run(hl, []int{})

	if len(res) != 1 {
		t.Fatalf("Expected 1 results, got %d instead\n", len(res))
	}

	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q instead\n", host, res[0].Host)
	}

	if !res[0].NotFound {
		t.Errorf("Expected host %q to be NOT found, got Found instead\n", host)
	}

	if len(res[0].PortStates) != 0 {
		t.Fatalf("Expected 0 port states, got %d instead\n", len(res[0].PortStates))
	}
}
