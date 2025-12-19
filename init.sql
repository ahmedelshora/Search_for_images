CREATE TABLE IF NOT EXISTS oc_product_description (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    language_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO oc_product_description (name, language_id) VALUES 
    ('CT17010 بلاور كهرباء 710 واط', 1),
    ('ام بى تى MED7503 شنيور عادة فقط 13 مم 750 وات اتجاهين', 1),
    ('ايه بى تى APT Dw02645 شنيور دقاق 13 مم 750 وات يمين وشمال سرعات', 1),
    ('كراون CT11023 راوتر حلية الكترونى متعدد الاستخدامات 710 وات', 1),
    ('دى سى ايه APB20C دريل فك و ربط 1/2 بوصة 340 وات + لقم', 1),
    ('دى سى ايه APB20C دريل فك و ربط 1/2 بوصة 340 وات + لقم', 1),
    ('بوش GSS 2300 صنفرة ترددية 190 x 93 مم 190 وات', 1);

