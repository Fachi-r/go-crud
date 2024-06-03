/**
 * Enter Receipt Number of purchased `#form`
      - The receipt number is used as a proof of payment for any of the forms.
      - The user selects what form they are signing up for (First year or returning student form) using  a `#checkbox` in this case
      - If they are applying for the first time, they only have to enter the receipt number. They are then presented with the first year form to enter all their details.  After the form has been submitted and their files uploaded successfully, their student record is created in the database with the given details. They are then redirected here with their new student loan number which will be displayed to them on this page.
        
      - If they are a returning student, they have to enter their student loan number as well, and the same procedure is carried out. The data of the student is then fetched accordingly to automatically fill in the form so that all they have to do is update their information where needed.
      - `#note` After someone applies, the receipt is removed from the database, to avoid someone trying to use the same receipt as proof of payment.
 */

// If we've been redirected after uploading files, we need to
// check url for upload details
const params = new URL(window.location.href).searchParams;
const success = params.get('success');
const loanNumber = params.get('loan_number');

if (success === "true") {
   // display Success modal
   document.getElementById("modal-btn").click()

   if (loanNumber == null) {
      document.querySelector(".modal-body").innerHTML =
         `<h4>Your details have been uploaded successfully!</h4>`
   } else {
      document.getElementById("loan-number-div").innerHTML = loanNumber
   }
}

const form = document.getElementById("form")
const checkbox = document.getElementById("checkbox")

// This runs on form submission
form.addEventListener("submit", validate)
// This runs everytime the checkbox is changed
checkbox.addEventListener('change', showOptions);

/** Show additional options for returning student login */
function showOptions() {
   var loanNumber = document.getElementById("loan_number")
   var options = document.querySelector(".options")

   // Disable or enable the input fields based on the checkbox value
   options.classList.toggle("enable", checkbox.checked)

   loanNumber.disabled = !checkbox.checked
   loanNumber.ariaDisabled = !checkbox.checked
}

/** Sends a request to the backend to check if the receipt number and/or student exists in the database*/
async function validate(event) {
   // prevent page from reloading
   event.preventDefault()

   console.log("Checking receipt...");

   // Validate Receipt number
   let receipt = document.getElementById("receipt_number")
   try {
      // Api returns JSON in the format {"exists": true || false}
      const response = await fetch(`http://localhost:5000/validate/receipts/${receipt.value}`);
      const data = await response.json();

      console.log("Result => ", data.exists);

      if (!data.exists) {
         receipt.classList.add("is-invalid")
         return
      }
   }
   catch (error) {
      console.error('Error:', error);
      return
   }

   // If returning student, validate student loan number and redirect to returning students form
   if (checkbox.checked) {
      console.log("Checking loan number...");

      let loanNumber = document.getElementById("loan_number")
      try {
         // Api returns JSON in the format {"exists": true || false}
         const response = await fetch(`http://localhost:5000/validate/students/${loanNumber.value}`);
         const data = await response.json();

         console.log("Result => ", data.exists);

         if (data.exists) {
            /* Redirect to returning student form with loan number and receipt*/
            window.location.href = `http://localhost:5000/forms/returning/?loan_number=${loanNumber.value}&receipt=${receipt.value}`
         } else {
            loanNumber.classList.add("is-invalid")
            return
         }
      } catch (error) {
         console.error('Error:', error);
         return
      }
   }
   else {
      /* Redirect to first year form with receipt number*/
      document.location.href = `http://localhost:5000/forms/first/?receipt=${receipt.value}`
   }
}