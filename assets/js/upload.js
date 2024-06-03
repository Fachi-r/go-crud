
// Get details from the URL
const params = new URL(window.location.href).searchParams;
const loanNumber = params.get('loan_number');
const formID = window.location.pathname.split("/")[2]

const submissionURL = `http://localhost:5000/api/forms/${formID}/docs?loan_number=${loanNumber}`;

const options = document.querySelector('.options');
// Disable or enable the input fields based on the form type
if (formID === "first") {
   options.classList.add("enable")
   const guardianNrc = document.getElementById('guardian_nrc');
   const tpin = document.getElementById('tpin');

   tpin.disabled = false
   tpin.ariaDisabled = false

   guardianNrc.disabled = false
   guardianNrc.ariaDisabled = false
}

const form = document.getElementById('form');
form.addEventListener('submit', (event) => {
   event.preventDefault()
   form.action = submissionURL
   form.submit()
})


