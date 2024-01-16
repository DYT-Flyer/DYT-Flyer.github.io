document.querySelectorAll('.circle').forEach(circle => {
  circle.addEventListener('click', function() {
    let valueToAdd = this.getAttribute('data-value'); // Corrected to use data-value
    let currentContent = document.getElementById('pn1').innerHTML;
    document.getElementById('pn1').innerHTML = currentContent + valueToAdd;
    console.log(valueToAdd);
  });
});

$(document).ready(function() {
  $('.tab-link').click(function() {
    // Get the target content ID from the data-target attribute
    let target = $(this).data('target');

    // Hide all tab content
    $('.tab-content').hide();

    // Show the target content
    $(target).show();
  });
});

function removeLastCharacter() {
    let currentContent = document.getElementById('pn1').innerHTML;
    document.getElementById('pn1').innerHTML = currentContent.slice(0,-1);
}

// Add event listener to the delete button
document.getElementById('db1').addEventListener('click', removeLastCharacter);

function addContact() {
  // Retrieve form data
  var name = document.getElementById('name').value;
  var phone = document.getElementById('phone').value;
  var email = document.getElementById('email').value;
  
  // Access the table body
  var tableBody = document.getElementById('address-book').getElementsByTagName('tbody')[0];

  // Create a new row
  var newRow = tableBody.insertRow();

  // Insert cells for name, phone, and email
  var nameCell = newRow.insertCell(0);
  var phoneCell = newRow.insertCell(1);
  var emailCell = newRow.insertCell(2);

  // Add the text to the cells
  nameCell.textContent = name;
  phoneCell.textContent = phone;
  emailCell.textContent = email;
  
  // Optionally clear the form fields after adding the contact
  document.getElementById('add-contact-form').reset();
}

document.getElementById('add-contact').addEventListener('click', addContact);

$(document).ready(function() {
  // Initially hide all tabs
  $('.tab-content').hide();
  
  // Show the first tab content by default
  $('.tab-content:first').show();
  $('.tab-link:first').addClass('active-tab');

  // Click event for tab links
  $('.tab-link').click(function() {
    // Remove active class from all tabs and hide all tab content
    $('.tab-link').removeClass('active-tab');
    $('.tab-content').hide();

    // Add active class to clicked tab and show corresponding tab content
    $(this).addClass('active-tab');
    $($(this).data('target')).show();
  });
});
