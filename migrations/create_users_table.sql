CREATE TABLE defect_reports (
                                id INT AUTO_INCREMENT PRIMARY KEY,
                                image_id INT NOT NULL,                -- 关联到 images 表的 id
                                verdict ENUM('OK', 'NG', 'UNKNOWN') DEFAULT 'UNKNOWN',  -- 检测结果
                                defects_count INT DEFAULT 0,          -- 检测到的缺陷数量
                                report_data JSON,                     -- 存储检测的详细数据，缺陷信息等，使用 JSON 格式
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
);

