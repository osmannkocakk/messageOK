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

### Some Working Environment Screen Shots During the Coding and Testing
- **DB Side**
  
![db screenshot](https://github.com/user-attachments/assets/d44ab9ea-5b9b-493f-b7fe-9f79fab1bba6)

- **Docker Side**

![docker images](https://github.com/user-attachments/assets/c05fd2d9-373f-4a73-b075-a21831958d95)

![docker container running](https://github.com/user-attachments/assets/6df8ed97-1b06-4521-8ba9-8b369d808ae6)

- **Swagger Side**

![Swagger Ana Ekran](https://github.com/user-attachments/assets/1ada19eb-734a-4faf-ae6c-62bd7f777056)
![Start Api Endpoint](https://github.com/user-attachments/assets/4004d75d-6d99-47db-998c-5a1e63131494)
![Stop Api Endpoint](https://github.com/user-attachments/assets/8b65a6e6-46e8-4302-a349-99737a03d231)
![Sent Messages](https://github.com/user-attachments/assets/af6f18d6-88b9-40c3-8436-1e6e6d839ba2)

- **Testing And Live Working Side**

![2 data](https://github.com/user-attachments/assets/ab43fd25-468c-4531-aff1-501ac7dda807)

![12de100](https://github.com/user-attachments/assets/1e9df8ec-6cba-4a8f-9e37-20fb348343a9)

![28de 100](https://github.com/user-attachments/assets/1e512ac9-7ea9-4742-bc60-775008766100)

![93de100](https://github.com/user-attachments/assets/11d34337-b6ef-4edd-ac20-08cc2390fd70)






