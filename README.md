# Message Service by Osman Koçak(-messageOK-)

This project is a messaging service written in Go with clean architecture principles. The service retrieves unsent 2 messages from a MySQL database and sends them to a specified endpoint every 2 minutes. Messages are cached in Redis once sent.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Project

1. **Clone the repository:**

   ```bash
   git clone https://github.com/osmannkocakk/messageOK.git
   cd messageOK
   ```

3. **Run Docker Compose:**

   ```bash
   cd docker
   docker-compose up --build
   ```

4. **Access Swagger Documentation:**

   Visit `http://localhost:8080/swagger/index.html` to view the Swagger documentation.

### API Endpoints

- **Start Automatic Sending Messages**: POST `/start`
- **Stop Automatic Sending Messages**: POST `/stop`
- **Get Sent Messages**: GET `/sent`

### Database Structure

```bash
CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content VARCHAR(160) NOT NULL,
    `to` VARCHAR(20) NOT NULL,
    status ENUM('unsent', 'sent') DEFAULT 'unsent',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_status (status),
    INDEX idx_to (`to`)
);
-- Some Sample data
INSERT INTO messages (content, `to`, status) VALUES
('Mutlu yillara Insider ailesi 1', '+905426460000', 'unsent'),
('Mutlu yillara Insider ailesi 2', '+905426460001', 'unsent'),
('Mutlu yillara Insider ailesi 3', '+905426460002', 'unsent');
 ```