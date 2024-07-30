let patientData = {
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
      { 
        date: "2023-05-15", 
        details: "Routine checkup - All vitals normal",
        diagnosis: "Healthy",
        treatment: "None required",
        prescription: "N/A",
        doctor: "Dr. Smith",
        cost: 100
      },
      { 
        date: "2022-11-03", 
        details: "Flu vaccination",
        diagnosis: "Preventive care",
        treatment: "Flu shot administered",
        prescription: "N/A",
        doctor: "Dr. Johnson",
        cost: 50
      },
      { 
        date: "2022-03-20", 
        details: "Sprained ankle",
        diagnosis: "Grade 1 ankle sprain",
        treatment: "RICE protocol (Rest, Ice, Compression, Elevation)",
        prescription: "Ibuprofen 400mg as needed for pain",
        doctor: "Dr. Brown",
        cost: 150
      }
    ]
  };
  
  function updatePatientInfo() {
    const infoTable = document.getElementById("patientInfoTable");
    infoTable.innerHTML = `
      <tr><th>Name</th><td>${patientData.name}</td></tr>
      <tr><th>ID</th><td>${patientData.id}</td></tr>
      <tr><th>Age</th><td>${patientData.age}</td></tr>
      <tr><th>Gender</th><td>${patientData.gender}</td></tr>
      <tr><th>Contact</th><td>${patientData.contact}</td></tr>
      <tr><th>Blood Group</th><td>${patientData.bloodGroup}</td></tr>
      <tr><th>Allergies</th><td>${patientData.allergies.join(", ")}</td></tr>
      <tr><th>Vaccines</th><td>${patientData.vaccines.join(", ")}</td></tr>
      <tr><th>Current Medications</th><td>${patientData.medications.join(", ")}</td></tr>
      <tr><th>Blood Pressure</th><td>${patientData.bloodPressure}</td></tr>
      <tr><th>Height</th><td>${patientData.height}</td></tr>
      <tr><th>Weight</th><td>${patientData.weight}</td></tr>
    `;
  }
  
  function updateSidebar() {
    const sidebar = document.getElementById("patientHistory");
    sidebar.innerHTML = "";
    const historyList = document.createElement("ul");
    patientData.records.forEach((record, index) => {
      const listItem = document.createElement("li");
      const link = document.createElement("a");
      link.textContent = `${record.date}: ${record.details.substring(0, 30)}...`;
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
  
  function updateMedicalRecords() {
    const recordsDiv = document.getElementById("medicalRecords");
    recordsDiv.innerHTML = "<h3>Recent Medical Records</h3>";
    const table = document.createElement("table");
    table.innerHTML = `
      <tr>
        <th>Date</th>
        <th>Details</th>
        <th>Doctor</th>
        <th>Cost</th>
        <th>Actions</th>
      </tr>
    `;
    patientData.records.forEach((record, index) => {
      const row = table.insertRow();
      row.innerHTML = `
        <td>${record.date}</td>
        <td>${record.details}</td>
        <td>${record.doctor}</td>
        <td>$${record.cost}</td>
        <td>
          <button class="btn" onclick="showFullHistory(${index})">View</button>
        </td>
      `;
    });
    recordsDiv.appendChild(table);
  }
  
  function showFullHistory(index) {
    const record = patientData.records[index];
    const modal = document.getElementById("historyModal");
    const content = document.getElementById("fullHistoryContent");
    content.innerHTML = `
      <h3>Date: ${record.date}</h3>
      <p><strong>Details:</strong> ${record.details}</p>
      <p><strong>Diagnosis:</strong> ${record.diagnosis}</p>
      <p><strong>Treatment:</strong> ${record.treatment}</p>
      <p><strong>Prescription:</strong> ${record.prescription}</p>
      <p><strong>Doctor:</strong> ${record.doctor}</p>
      <p><strong>Cost:</strong> $${record.cost}</p>
    `;
    modal.style.display = "block";
  }
  
  function downloadRecord() {
    const content = document.getElementById("fullHistoryContent").innerText;
    const blob = new Blob([content], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = "medical_record.txt";
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  }
  
  function emailRecord() {
    const email = document.getElementById("emailInput").value;
    if (!email) {
      alert("Please enter a valid email address.");
      return;
    }
    // In a real application, you would send this to your server to handle the email sending
    alert(`Medical record will be sent to ${email}. (This is a simulation)`);
  }
  
  function logout() {
    if (confirm("Are you sure you want to log out?")) {
      alert("You have been logged out successfully.");
      // Here you would typically redirect to a login page or perform other logout actions
      // For this example, we'll just reload the page
      location.reload();
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
  
  // Initialize the dashboard
  updatePatientInfo();
  updateSidebar();
  updateMedicalRecords();