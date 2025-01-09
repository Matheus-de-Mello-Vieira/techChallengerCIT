const API_URL = "http://localhost:8081";

document.addEventListener("DOMContentLoaded", async () => {
  const totals = await getTotals();
  displayTotal(totals["GeneralTotal"]);
  displayChartByHour(totals["TotalByHour"]);
  displayChartByParticipant(totals["TotalByParticipant"]);
});

async function getTotals(retries = 3) {
  try {
    const response = await fetch(`${API_URL}/votes/thorough`);

    if (!response.ok) {
      throw new Error("Falha ao obter resultados");
    }
    return await response.json();
  } catch (error) {
    console.error("Erro ao obter resultados:", error);
    if (retries > 0) {
      console.log(`Tentando novamente. Tentativas restantes: ${retries - 1}`);
      setTimeout(() => getResults(retries - 1), 2000);
    } else {
      alert(
        "Ocorreu um erro ao tentar obter os resultados. Por favor, tente novamente mais tarde."
      );
    }
  }
}

function displayTotal(total) {
  const span = document.getElementById("generalTotal");

  span.innerHTML = total;
}

function displayChartByHour(totalByHour) {
  const hours = totalByHour.map((x) => x["Hour"]);
  const totals = totalByHour.map((x) => x["Total"]);

  const data = {
    labels: hours,
    datasets: [
      {
        label: "Totais por hora",
        data: totals,
        fill: false,
        borderColor: "rgb(75, 192, 192)",
        tension: 0.1,
      },
    ],
  };

  const config = {
    type: "line",
    data: data,
  };

  // 4. Create and render the chart
  const ctx = document.getElementById("hourlyTotalChart").getContext("2d");
  const resultChart = new Chart(ctx, config);
}

function displayChartByParticipant(TotalByParticipant) {
  const labels = Object.keys(TotalByParticipant);
  const totals = Object.values(TotalByParticipant);

  const data = {
    labels: labels,
    datasets: [
      {
        label: "Totais por participante",
        data: totals,
        backgroundColor: ["rgba(54, 162, 235, 0.2)"],
        borderColor: ["rgba(54, 162, 235, 1)"],
        borderWidth: 1,
      },
    ],
  };

  const config = {
    type: "bar",
    data: data,
    options: {
      scales: {
        y: {
          beginAtZero: true,
        },
      },
    },
  };

  const ctx = document.getElementById("chartbyParticipant").getContext("2d");
  const resultChart = new Chart(ctx, config);
}