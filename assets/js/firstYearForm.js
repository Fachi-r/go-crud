/**
 * First Year Form:
 * - This page is reached only through the Login page. When reached, URL parameters for the receipt number are accessed
 * and displayed on the page under the appropriate fields
 * 
 * - Dynamic fields for the Agreement Form, such as the `Date`, are also updated on page load
 * - The `YearOfStudy` field is automatically set to first year, and disabled.
 * - After the user enters all their details and submits, They are Sent to another page which displays Their newly assigned loan number. This will be used on their next student loan renewal as a returning student.
 * 
 * Process:   
 * The student details are loaded into this script when the form is submitted
   The data is added to a Student Object Model resembling the same one used in the database.
   The data is parsed into `JSON` and sent in a `POST` request to the backend, which then creates a user model in the database and sends back the data that it has uploaded
 *
 * From this returned Data, the Student Loan Number is gotten and displayed to the user

   `#note` - The Student and Guardian are two separate tables. But in this form, they are sent in a single object. The backend separates the fields appropriately and handles any errors
 */
const form = document.getElementById("form")
form.addEventListener("submit", submitForm)

// Loan number is automatically generated for first year students
var LoanNumber = Math.floor(Math.random() * 10000000000)

var formIDDiv = document.querySelector(".form_id")
formIDDiv.innerText = Math.floor(Math.random() * 10000)

// Get Data from URL and update the appropriate fields.
// Assuming "localhost:5000/forms/?receipt=1234" is your URL
var params = new URLSearchParams(window.location.search)
var receipt = params.get('receipt'); // is the string "1234"
var receiptDiv = document.querySelector(".receipt_number")
receiptDiv.innerHTML = `${receipt}`

// Update all dates to today's date
var dateDivs = document.querySelectorAll(".date")
dateDivs.forEach(date => {
   date.innerHTML = `${new Date().toDateString()}`
});

// Update all name fields to the student name
var NameDiv = document.getElementById("student-name")
var studentNameDivs = document.querySelectorAll(".student_name")
NameDiv.addEventListener("change", (event) => {
   studentNameDivs.forEach(div => {
      div.innerHTML = event.target.value
   })
})

/** This function gets all the data from the form inputs, and creates an Object with them to post to the database */
async function submitForm(event) {
   // Prevent page from reloading
   event.preventDefault()
   // Get Values from Form
   let Programme = document.getElementById("student-program").value
   let YearOfStudy = document.getElementById("student-year").value
   let StudentNumber = document.getElementById("student-number").value
   let NRC = document.getElementById("student-nrc").value

   let Degree = document.getElementById("student-degree").value
   let School = document.getElementById("student-school").value

   let Bank = document.getElementById("bank-name").value
   let Branch = document.getElementById("bank-branch").value
   let AccountName = document.getElementById("account-name").value
   let AccountNumber = document.getElementById("account-number").value

   let GuardianName = document.getElementById("guardian-name").value
   let GuardianNRC = document.getElementById("guardian-nrc").value
   let Nationality = document.getElementById("guardian-nationality").value
   var Gender = document.getElementById("guardian-gender").value;
   let Relationship = document.getElementById("guardian-relationship").value
   let Address = document.getElementById("guardian-address").value
   let Town = document.getElementById("guardian-town").value
   let Province = document.getElementById("guardian-province").value
   let PostalAddress = document.getElementById("guardian-postal_address").value
   let Phone = document.getElementById("guardian-phone").value
   let Email = document.getElementById("guardian-email").value

   // Loan number is automatically generated for first year students
   let LoanNumber = Math.floor(Math.random() * 10000000000)

   // Make sure all keys match the keys used in the database models exactly
   var formData = {
      // Student Details
      LoanNumber: LoanNumber,
      NRC: NRC,
      Name: NameDiv.value,
      Programme: Programme,
      Degree: Degree,
      School: School,
      YearOfStudy: Number(YearOfStudy),
      StudentNumber: Number(StudentNumber),
      // Bank Details
      Bank: Bank,
      Branch: Branch,
      AccountName: AccountName,
      AccountNumber: Number(AccountNumber),
      // Guardian Details
      GuardianName: GuardianName,
      GuardianNRC: GuardianNRC,
      Relationship: Relationship,
      Gender: Gender,
      Nationality: Nationality,
      Address: Address,
      Town: Town,
      Province: Province,
      PostalAddress: PostalAddress,
      Phone: Phone,
      Email: Email
   }

   console.log("Posting data...");

   console.log("Student Number => ", formData.StudentNumber);
   try {
      const response = await fetch("http://localhost:5000/api/students", {
         method: 'POST',
         headers: {
            'Content-type': 'application/json; charset=UTF-8'
         },
         body: JSON.stringify(formData)
      });

      if (!response.ok) {
         throw new Error(`HTTP error! status: ${response.status}`);
      }

      let data = await response.json();
      console.log(data);

      // After Form Submission, show Student Loan Number
      let body = document.querySelector(".body")
      body.innerHTML = `<code>${data}</code>`
   } catch (error) {
      console.error('Error:', error);
   }
}
