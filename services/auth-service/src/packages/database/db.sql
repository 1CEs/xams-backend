USE xams;

CREATE TABLE IF NOT EXISTS users(
    user_id VARCHAR(20) PRIMARY KEY UNIQUE NOT NULL,
    personal_id VARCHAR(13) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    prename VARCHAR(50) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    branch_id INT NOT NULL,
    role ENUM('student', 'teacher', 'admin') NOT NULL,
    FOREIGN KEY(branch_id) REFERENCES branch(branch_id)
);

CREATE TABLE IF NOT EXISTS faculty (
    faculty_id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    faculty_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS branch (
    branch_id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    faculty_id INT NOT NULL,
    branch_name VARCHAR(255) NOT NULL UNIQUE,
    FOREIGN KEY(faculty_id) REFERENCES faculty(faculty_id)
);