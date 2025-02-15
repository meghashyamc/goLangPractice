package test

import(
"testing"
"store"
"delete"

)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee


func TestDeleteFullEmployeeList(t *testing.T){

initializeStuff()
if len(testEmployees) == 0 {

	t.Error("Problem in initializing delete tests.")
}
delete.DeleteFullEmployeeList(&testEmployees)

if len(testEmployees) != 0 {

	t.Error("Expected all employees to be deleted but found", len(testEmployees), "left")
}

}


func TestDeleteFullIDempMap(t *testing.T){

	initializeStuff()
	if len(testEmployees) == 0 {
	
		t.Error("Problem in initializing delete tests.")
	}
	
	delete.DeleteFullIDempMap(&(store.IdEmpMap))

	if len(store.IdEmpMap) != 0 {
	
		t.Error("Expected all employees to be deleted but found", len(store.IdEmpMap), "left")
	}
	
	}


	func TestDeleteFullDeptEmpMap(t *testing.T){

		initializeStuff()
		if len(testEmployees) == 0 {
		
			t.Error("Problem in initializing delete tests.")
		}
		
		delete.DeleteFullDeptEmpMap(&(store.DeptEmpMap))
	
		if len(store.DeptEmpMap) != 0 {
		
			t.Error("Expected all employees to be deleted but found", len(store.DeptEmpMap), "left")
		}
		
		}

		func TestDeleteFullLocEmpMap(t *testing.T){

			initializeStuff()
			if len(testEmployees) == 0 {
			
				t.Error("Problem in initializing delete tests.")
			}
			
			delete.DeleteFullLocEmpMap(&(store.LocEmpMap))
		
			if len(store.LocEmpMap) != 0 {
			
				t.Error("Expected all employees to be deleted but found", len(store.LocEmpMap), "left")
			}
			
			}

func TestDeleteEmployeeByIDFromList(t *testing.T){

initializeStuff()

if len(testEmployees) != 3{
	t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
}

if (testEmployees[0]).Name != "Pappu"{
	t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
}

delete.DeleteEmployeeByIDFromList(1, &testEmployees)

if len(testEmployees) != 2{
	t.Error("Expected size of testEmployees to be 2 after 1 deletion got", len(testEmployees))
}

if (testEmployees[0]).Name == "Pappu"{
	t.Error("Expected Pappu to be deleted got",(testEmployees[0]).Name)}

}


func TestDeleteEmployeeByIDFromDeptMap(t *testing.T){

	initializeStuff()
	
	if len(testEmployees) != 3{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	if (testEmployees[0]).Name != "Pappu"{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	delete.DeleteEmployeeByIDFromDeptMap(1, testEmployees[0].GetDept(),&(store.DeptEmpMap))
	
	for _, empl := range *(store.DeptEmpMap["Accounts"]){

		if empl.GetID() == 1{

			t.Error("Expected ID 1 to be deleted from Accounts but found", empl.GetID())
		}

	}

	for _, empl := range *(store.DeptEmpMap["Management"]){

		if empl.GetID() == 1{

			t.Error(store.DeptEmpMap["Management"])

			t.Error("Expected ID 1 to be deleted from Management but found", empl.GetID())
		}

	}
	
}


func TestDeleteEmployeeByIDFromLocMap(t *testing.T){

	initializeStuff()
	
	if len(testEmployees) != 3{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	if (testEmployees[0]).Name != "Pappu"{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	delete.DeleteEmployeeByIDFromLocMap(1, testEmployees[0].GetPins(),&(store.LocEmpMap))
	
	for _, empl := range *(store.LocEmpMap[560002]){

		if empl.GetID() == 1{

			t.Error("Expected ID 1 to be deleted from 560002 but found", empl.GetID())
		}

	}

	for _, empl := range *(store.LocEmpMap[560103]){

		if empl.GetID() == 1{

			t.Error(store.LocEmpMap[560103])

			t.Error("Expected ID 1 to be deleted from 560103 but found", empl.GetID())
		}

	}
	
}


func TestDeleteEmployeeByIDFromIDMap(t *testing.T){

	initializeStuff()
	
	if len(testEmployees) != 3{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	if (testEmployees[0]).Name != "Pappu"{
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}
	
	delete.DeleteEmployeeByIDFromIdMap(1,&(store.IdEmpMap))
	
	
	if(store.IdEmpMap[1]).Name == "Pappu"{

			t.Error("Expected Pappu to be deleted from ID Emp map but found", (store.IdEmpMap[1]).Name)}
		


	
}

func initializeStuff(){
	
	address1 := store.Address{Doorno: 25,
		Street: "Adepalli Street",
		Locality: "Basvangudi",
		PIN: 560002}
		address2 := store.Address{Doorno: 32,
			Street: "Pilli Street",
			Locality: "Balagere",
			PIN: 560103}
		
		address3 := store.Address{Doorno: 564,
			Street: "White House Street",
			Locality: "Malleshwaram",
			PIN: 560003}
		
		address4 := store.Address{Doorno: 43,
			Street: "Vishveshwarya Street",
			Locality: "Hosur",
			PIN: 560203}
			testEmployee1 = store.Employee{
		EmpID: 1,
		Name: "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses: [](store.Address){address1, address2},
		There: true,
		
			}
		
			testEmployee2 = store.Employee{
				EmpID: 2,
				Name: "Rinky",
				Department: []string{"IT", "Admin"},
				Addresses: [](store.Address){address2, address3},
				There: true,
				
					}
		
					testEmployee3 = store.Employee{
						EmpID: 3,
						Name: "Tinkadi",
						Department: []string{"Accounts", "Admin"},
						Addresses: [](store.Address){address4},
						There: true,
						
							}
			
			 testEmployees = []store.Employee {
				testEmployee1, testEmployee2, testEmployee3}
		
		store.StoreEmployeesByIdDeptAndLoc(&testEmployees)


}