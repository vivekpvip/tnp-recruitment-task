document.getElementById("rsvpForm").addEventListener("submit", function(event) {
    event.preventDefault();
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const phone = document.getElementById("phone").value;
    const college = document.getElementById("college").value;
    const year = document.getElementById("year").value;
    const expectations = document.getElementById("expectations").value;

    if (!name || !email || !phone || !college || !year || !expectations) {
        document.getElementById("responseMessage").innerText = "❌ Please fill all fields!";
        return;
    }

    localStorage.setItem("registered", JSON.stringify({ name, email, phone, college, year, expectations, checkedIn: false }));
    document.getElementById("responseMessage").innerText = "✅ Registration Successful!";
});

document.getElementById("checkinBtn").addEventListener("click", function() {
    let user = JSON.parse(localStorage.getItem("registered"));
    if (user && !user.checkedIn) {
        user.checkedIn = true;
        localStorage.setItem("registered", JSON.stringify(user));
        document.getElementById("statusMessage").innerText = `✅ ${user.name}, You have checked in!`;
    } else if (!user) {
        document.getElementById("statusMessage").innerText = "❌ You need to register first!";
    } else {
        document.getElementById("statusMessage").innerText = "❌ You have already checked in!";
    }
});

document.getElementById("checkoutBtn").addEventListener("click", function() {
    let user = JSON.parse(localStorage.getItem("registered"));
    if (user && user.checkedIn) {
        user.checkedIn = false;
        localStorage.setItem("registered", JSON.stringify(user));
        document.getElementById("statusMessage").innerText = `✅ ${user.name}, You have checked out!`;
    } else {
        document.getElementById("statusMessage").innerText = "❌ You need to check-in first!";
    }
});
