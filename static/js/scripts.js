document.addEventListener('DOMContentLoaded', function() {
    const features = document.querySelectorAll('.feature');
    let currentFeatureIndex = 0;

    function showFeature(index) {
        features.forEach((feature, i) => {
            feature.classList.toggle('active', i === index);
        });
    }

    function nextFeature() {
        currentFeatureIndex = (currentFeatureIndex + 1) % features.length;
        showFeature(currentFeatureIndex);
    }

    showFeature(currentFeatureIndex);
    setInterval(nextFeature, 5000);
});


// dashboard facility

let currentPatient = null;
let generatedOTP = null;

function searchPatient() {
  // Simulating a patient search
  currentPatient = {
    name: "John Doe",
    id: "12345",
    age: 35,
    gender: "Male",
    contact: "+1234567890",
    bloodGroup: "A+",
    allergies: ["Penicillin", "Peanuts"],
    vaccines: ["COVID-19", "Flu", "Tetanus"],
    hivStatus: "Negative",
    medications: ["Lisinopril", "Metformin"],
    bloodPressure: "120/80",
    height: "180 cm",
    weight: "75 kg",
    records: [
      { date: "2023-05-15", visitType: "Routine Checkup", details: "All vitals normal", cost: 100 },
      { date: "2022-11-03", visitType: "Vaccination", details: "Flu vaccination administered", cost: 50 },
      { date: "2022-03-20", visitType: "Emergency", details: "Sprained ankle - Prescribed rest and ice", cost: 150 },
      { date: "2021-09-10", visitType: "Specialist Consultation", details: "Cardiology follow-up - EKG performed", cost: 200 },
      { date: "2021-06-05", visitType: "Routine Checkup", details: "Annual physical - All tests within normal range", cost: 120 }
    ]
  };
  updatePatientInfo();
  document.getElementById("otpVerification").style.display = "block";
  generatedOTP = Math.floor(100000 + Math.random() * 900000).toString();
  alert(`OTP sent to patient's contact: ${generatedOTP}`);
}

function updatePatientInfo() {
  const infoTable = document.getElementById("patientInfoTable");
  infoTable.innerHTML = `
    <tr><th>Name</th><td>${currentPatient.name}</td></tr>
    <tr><th>ID</th><td>${currentPatient.id}</td></tr>
    <tr><th>Age</th><td>${currentPatient.age}</td></tr>
    <tr><th>Gender</th><td>${currentPatient.gender}</td></tr>
    <tr><th>Contact</th><td>${currentPatient.contact}</td></tr>
  `;
  document.getElementById("medicalRecords").innerHTML = "<p>Verify OTP to view full medical records.</p>";
}

function verifyOTP() {
  const enteredOTP = document.getElementById("otpInput").value;
  if (enteredOTP === generatedOTP) {
    alert("OTP verified. Access granted to full medical records.");
    displayFullMedicalRecords();
    document.getElementById("otpVerification").style.display = "none";
    document.getElementById("newRecordForm").style.display = "block";
  } else {
    alert("Invalid OTP. Please try again.");
  }
}

function displayFullMedicalRecords() {
  const recordsHTML = `
    <h3>Full Medical Records</h3>
    <table>
      <tr><th>Blood Group</th><td>${currentPatient.bloodGroup}</td></tr>
      <tr><th>Allergies</th><td>${currentPatient.allergies.join(", ")}</td></tr>
      <tr><th>Vaccines</th><td>${currentPatient.vaccines.join(", ")}</td></tr>
      <tr><th>HIV Status</th><td>${currentPatient.hivStatus}</td></tr>
      <tr><th>Medications</th><td>${currentPatient.medications.join(", ")}</td></tr>
      <tr><th>Blood Pressure</th><td>${currentPatient.bloodPressure}</td></tr>
      <tr><th>Height</th><td>${currentPatient.height}</td></tr>
      <tr><th>Weight</th><td>${currentPatient.weight}</td></tr>
    </table>
  `;
  document.getElementById("medicalRecords").innerHTML = recordsHTML;
  updateSidebar();
}

function updateSidebar() {
  const sidebar = document.getElementById("patientHistory");
  sidebar.innerHTML = "<h3>Recent Medical History</h3>";
  const historyList = document.createElement("ul");
  currentPatient.records.slice(0, 3).forEach((record, index) => {
    const listItem = document.createElement("li");
    const link = document.createElement("a");
    link.textContent = `${record.date}: ${record.visitType}`;
    link.href = "#";
    link.className = "history-link";
    link.onclick = (e) => {
      e.preventDefault();
      showFullHistory(index);
    };
    listItem.appendChild(link);
    historyList.appendChild(listItem);
  });
  sidebar.appendChild(historyList);
}

function showFullHistory(index) {
  const record = currentPatient.records[index];
  const modal = document.getElementById("historyModal");
  const content = document.getElementById("fullHistoryContent");
  content.innerHTML = `
    <h3>Date: ${record.date}</h3>
    <p><strong>Visit Type:</strong> ${record.visitType}</p>
    <p><strong>Details:</strong> ${record.details}</p>
    <p><strong>Cost:</strong> $${record.cost}</p>
  `;
  modal.style.display = "block";
}

function showAllHistory() {
  const modal = document.getElementById("historyModal");
  const content = document.getElementById("fullHistoryContent");
  content.innerHTML = "<h2>All Medical History</h2>";
  const table = document.createElement("table");
  table.innerHTML = `
    <tr>
      <th>Date</th>
      <th>Visit Type</th>
      <th>Details</th>
      <th>Cost</th>
    </tr>
  `;
  currentPatient.records.forEach(record => {
    const row = table.insertRow();
    row.innerHTML = `
      <td>${record.date}</td>
      <td>${record.visitType}</td>
      <td>${record.details}</td>
      <td>$${record.cost}</td>
    `;
  });
  content.appendChild(table);
  modal.style.display = "block";
}

function submitNewRecord() {
  const visitDate = document.getElementById("visitDate").value;
  const visitType = document.getElementById("visitType").value;
  const symptoms = document.getElementById("symptoms").value;
  const diagnosis = document.getElementById("diagnosis").value;
  const treatment = document.getElementById("treatment").value;
  const prescription = document.getElementById("prescription").value;
  const notes = document.getElementById("notes").value;
  const cost = document.getElementById("cost").value;
  const files = document.getElementById("fileUpload").files;

  if (!visitDate || !visitType || !cost) {
    alert("Please fill in all required fields (Visit Date, Visit Type, and Cost).");
    return;
  }

  const newRecord = {
    date: visitDate,
    visitType: visitType,
    details: `Symptoms: ${symptoms}\nDiagnosis: ${diagnosis}\nTreatment: ${treatment}\nPrescription: ${prescription}\nNotes: ${notes}`,
    cost: parseFloat(cost),
    files: Array.from(files).map(file => file.name)
  };

  currentPatient.records.unshift(newRecord);
  alert("New medical record added successfully!");
  updateSidebar();
  document.getElementById("medicalRecordForm").reset();
  document.getElementById("fileList").textContent = "";
}

function logout() {
    if (confirm("Are you sure you want to log out?")) {
  
      alert("You have been logged out successfully.");
      // Redirect to the login page
      window.location.href = '/login'; // Update with your actual login page URL
    }
  }
  

// Modal close functionality
const modal = document.getElementById("historyModal");
const span = document.getElementsByClassName("close")[0];
span.onclick = function() {
  modal.style.display = "none";
}
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
}

// File upload handling
document.getElementById("fileUpload").addEventListener("change", function(event) {
  const fileList = event.target.files;
  let fileNames = "";
  for (let i = 0; i < fileList.length; i++) {
    fileNames += fileList[i].name + ", ";
  }
  document.getElementById("fileList").textContent = fileNames.slice(0, -2);
});

// Initialize the dashboard
searchPatient();
