/**
 * Returning Student Form:
 * - This page is reached only through the Login page. When reached, URL parameters for the receipt number and loan number are accessed and displayed on the page under the appropriate fields
 * 
 * - On page load, all student details with the provided loan number are fetched from the database.
 * - The appropriate fields are updated into the form automatically
 * - Specific fields are disabled to prevent the student from updating info that shouldn't change
 * - After the user Edits all their details and submits, they're are sent to a success page if details were updated successfully
 * 
 * Process:   
 * - The student details are loaded into this script when the form is submitted
   - The data is added to a Student Object Model resembling the same one used in the database.
   - The data is parsed into `JSON` and sent in a `PUT` request to the backend, which 
   then updates the student details and sends back a success flag

   `#note` - The Student and Guardian are two separate tables. That's why in this form, they are fetched separately one after the other.
 */


/* 
1. Asign all form fields to variables 
*/
const YearOfStudy = document.getElementById("student-year")
const Programme = document.getElementById("student-program")
const LoanNumberDiv = document.getElementById("student-loan_number")
const StudentNumber = document.getElementById("student-number")
const NRC = document.getElementById("student-nrc")

const Name = document.getElementById("student-name")
const Degree = document.getElementById("student-degree")
const School = document.getElementById("student-school")

const Bank = document.getElementById("bank-name")
const Branch = document.getElementById("bank-branch")
const AccountName = document.getElementById("account-name")
const AccountNumber = document.getElementById("account-number")

const GuardianName = document.getElementById("guardian-name")
const GuardianNRC = document.getElementById("guardian-nrc")
const Nationality = document.getElementById("guardian-nationality")
const Gender = document.getElementById("guardian-gender")
const Relationship = document.getElementById("guardian-relationship")
const Address = document.getElementById("guardian-address")
const Town = document.getElementById("guardian-town")
const Province = document.getElementById("guardian-province")
const PostalAddress = document.getElementById("guardian-postal_address")
const Phone = document.getElementById("guardian-phone")
const Email = document.getElementById("guardian-email")

/* 
2. Update static fields 
*/
// Get Data from URL and update the appropriate fields.
// Assuming "localhost:5000/forms/returning?receipt=1234&loan_number=9876" is your URL
const params = new URLSearchParams(window.location.search)
const receipt = Number(params.get('receipt')) // is the string "1234"
const loan_number = Number(params.get("loan_number"))
const form = document.getElementById("form")
const formIDDiv = document.querySelector(".form_id")
const dateDivs = document.querySelectorAll(".date")
const studentNameDivs = document.querySelectorAll(".student_name")

// Update form ID with random ID
formIDDiv.innerText = Math.floor(Math.random() * 10000)

// Update Form Number based on year
YearOfStudy.addEventListener("change", (event) => {
   document.getElementById("form_number").innerHTML = "Form " + event.target.value
})

// Update all dates to today's date
dateDivs.forEach(date => {
   date.innerHTML = `${new Date().toDateString()}`
})

/* 
3. Retrieve student data using loan_number
*/
var student, guardian
// Get student from Database
async function getStudent() {
   try {
      let response = await fetch(`http://localhost:5000/api/students/${loan_number}`)
      if (!response.ok) {
         throw new Error("Error: Couldn't fetch Student: ", response.status)
      }
      let data = await response.json()
      return data.data
   } catch (error) {
      console.error(error)
   }
}
// Get guardian with ID from database
async function getGuardian(id) {
   try {
      let response = await fetch(`http://localhost:5000/api/guardians/${id}`)
      if (!response.ok) {
         throw new Error("Error: Couldn't fetch Guardian: ", response.status)
      }
      let data = await response.json()
      return data.data
   } catch (error) {
      console.error(error)
   }
}

/*
4. Update Form Data
*/
async function updateForm() {
   student = await getStudent()
   guardian = await getGuardian(student.Guardian)

   if (guardian != undefined && student != undefined) {
      Programme.value = student.Programme
      YearOfStudy.value = student.YearOfStudy
      LoanNumberDiv.value = student.LoanNumber
      StudentNumber.value = student.StudentNumber
      NRC.value = student.NRC

      Name.value = student.Name
      Degree.value = student.Degree
      School.value = student.School

      Bank.value = student.Bank
      Branch.value = student.Branch
      AccountName.value = student.AccountName
      AccountNumber.value = student.AccountNumber

      GuardianName.value = guardian.GuardianName
      GuardianNRC.value = guardian.GuardianNRC
      Nationality.value = guardian.Nationality
      Gender.value = guardian.Gender
      Relationship.value = guardian.Relationship
      Address.value = guardian.Address
      Town.value = guardian.Town
      Province.value = guardian.Province
      PostalAddress.value = guardian.PostalAddress
      Phone.value = guardian.Phone
      Email.value = guardian.Email
   }

   // Update all name fields to the student name
   studentNameDivs.forEach(div => {
      div.innerHTML = student.Name
   })
}

updateForm()

/** 
 * This function gets all the data from the form inputs, and creates an Object with them to post to the database 
 * */
async function submitForm() {
   // Make sure all keys match the keys used in the database models exactly
   let studentData = {
      YearOfStudy: Number(YearOfStudy.value),
      Bank: Bank.value,
      Branch: Branch.value,
      AccountName: AccountName.value,
      AccountNumber: Number(AccountNumber.value),
   }
   let guardianData = {
      GuardianName: GuardianName.value,
      GuardianNRC: GuardianNRC.value,
      Nationality: Nationality.value,
      Gender: Gender.value,
      Relationship: Relationship.value,
      Address: Address.value,
      Town: Town.value,
      Province: Province.value,
      PostalAddress: PostalAddress.value,
      Phone: Phone.value,
      Email: Email.value
   }

   console.log("Updating student data...");
   try {
      const response = await fetch(`http://localhost:5000/api/students/${student.LoanNumber}`, {
         method: 'PUT',
         headers: {
            'Content-type': 'application/json; charset=UTF-8'
         },
         body: JSON.stringify(studentData)
      });

      if (!response.ok) {
         throw new Error(`HTTP error! status: ${response.status}`);
      }

      let data = await response.json();
      console.log(data);

   } catch (error) {
      console.error('Failed to update Student:', error);
   }

   console.log("Updating guardian data...");
   try {
      const response = await fetch(`http://localhost:5000/api/guardians/${student.Guardian}`, {
         method: 'PUT',
         headers: {
            'Content-type': 'application/json; charset=UTF-8'
         },
         body: JSON.stringify(guardianData)
      });

      if (!response.ok) {
         throw new Error(`HTTP error! status: ${response.status}`)
      }

      let data = await response.json();
      console.log(data);

   } catch (error) {
      console.error("Failed to update Guardian: ", error)
   }
}

async function deleteReceipt() {
   try {
      const response = await fetch(`http://localhost:5000/api/receipts/${receipt}`, {
         method: 'DELETE'
      });
      if (!response.ok) {
         throw new Error(`Failed to delete Receipt: ${response.status}`)
      }
   } catch (error) {
      console.error('Error:', error)
   }
}

form.addEventListener("submit", async (event) => {
   event.preventDefault()
   await submitForm()
   await deleteReceipt()
   // Redirect to document upload page with loan number
   window.location.href = `http://localhost:5000/forms/returning/docs?loan_number=${loan_number}`
})
