
# Transaction Processor Microservice

This project is a microservice developed in Go that processes credit and debit transactions from a CSV file. The service calculates a transaction summary, sends an email with the processed information, and stores the transactions in a SQL Server database.

---

## 📋 Main Features

- Read CSV files containing transactions.
- Calculate financial summary:
  - Total balance.
  - Number of transactions per month.
  - Average credit and debit amounts per month.
- Send an email with the generated summary.
- Persist transactions into a MSSQL database.
- HTTP API with a single endpoint `POST /process`.

---

## 🛠️ Technologies Used

- **Golang** (v1.23)
- **Gin** for the HTTP server
- **Gomail** for sending emails
- **Microsoft SQL Server** for persistence
- **Docker** and **Docker Compose** for container orchestration

---

## ⚙️ Installation and Execution

### 1. Clone the repository

```bash
git clone <your-repo>
cd transaction_processor
```

### 2. Configure environment variables

Modify the `config.env` file with your information:

```env
FILE_DIRECTORY=./data
PORT=8080

# MSSQL Database
MSSQL_HOST=mssql_server
MSSQL_USER=SA
MSSQL_PASSWORD=Admin_2024
MSSQL_PORT=1433
MSSQL_NAME=Stori

# Email SMTP Config
SMTP_HOST=smtp.gmail.com
SMTP_PORT=465
SENDER_EMAIL=your_email@gmail.com
RECIPIENT_EMAIL=recipient_email@example.com
SMTP_USERNAME=your_email@gmail.com
SMTP_PASS=your_application_password
```

> 📀 Remember to use a [Gmail application password](https://support.google.com/accounts/answer/185833) if you are using a Gmail account.

---

### 3. Start the project with Docker Compose

```bash
docker-compose up --build
```

This will:

- Build the service image.
- Create the necessary containers:
  - `transaction_processor` (the Go microservice)
  - `mssql_server` (the SQL Server instance)
  - `db_init` (for initializing the `Stori` database).

✅ The server will listen on `http://localhost:8080`.

---

## 🚀 How to Use the Service

Once running, you can process transactions by sending a `POST` request:

### Using `curl`:

```bash
curl -X POST http://localhost:8080/process
```

### Using Postman:

- Method: `POST`
- URL: `http://localhost:8080/process`
- Body: (empty)

The service will:

- Read the CSV file located at `./data/transactions.csv`
- Process the financial summary
- Send an email to the configured recipient
- Save the transactions into the `transactions` table of the database.

---

## 📂 Project Structure

```
transaction_processor/
├── commands/              # Commands representing each processing step
├── core/                  # Common data structures (context, summary, transaction)
├── facade/                # Facade to orchestrate the entire flow
├── utils/                 # Utilities: MSSQL connection, config loading
├── data/                  # CSV file (transactions.csv)
├── Dockerfile             # Application Dockerfile
├── docker-compose.yml     # Container orchestration file
├── config.env             # Environment variables
├── go.mod / go.sum        # Go module dependencies
└── main.go                # Entry point
```

---

## 🛠️ Important Technical Notes

- **Persistence in MSSQL**: The service stores each transaction (date, amount) into the `transactions` table in the `Stori` database.
- **Error handling**: If any error occurs (reading the file, sending the email, saving to the database), the service responds with an appropriate JSON error.
- **Docker Volumes**: The `data` directory is mounted into the container to access the `transactions.csv` file.

---

## 🧰 Future Improvements (Optional)

- Push logs to an observability system (like ELK Stack or Grafana Loki).
- Expose Prometheus metrics from the microservice.
- Allow dynamic CSV upload via `POST` (multipart/form-data).
- Implement basic authentication to secure the `/process` endpoint.

---

## 📜 License

Technical challenge project. Free to use for educational and evaluation purposes.

---
