// Copyright 2023 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoorsdk

import (
	"testing"
	"time"
)

func TestPayment(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Payment")

	// Add a new object
	payment := &Payment{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  time.Now().Format(time.RFC3339),
		DisplayName:  name,
		ProductName:  "casbin",
	}
	_, err := AddPayment(payment)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	payments, err := GetPayments()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range payments {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	payment, err = GetPayment(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if payment.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", payment.Name, name)
	}

	// Update the object
	updatedProductName := "Updated Casdoor Website"
	payment.ProductName = updatedProductName
	_, err = UpdatePayment(payment)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedPayment, err := GetPayment(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedPayment.ProductName != updatedProductName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedPayment.ProductName, updatedProductName)
	}

	// Delete the object
	_, err = DeletePayment(payment)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedPayment, err := GetPayment(name)
	if err != nil || deletedPayment != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
