use loker;

CREATE TABLE employment_types (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE position_levels (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    position_level VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE jobs (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    position_level_id INT NOT NULL,
    employment_type_id INT NOT NULL,
    status BOOL NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_position_level FOREIGN KEY (position_level_id)
        REFERENCES position_levels (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_employment_type FOREIGN KEY (employment_type_id)
        REFERENCES employment_types (id)
        ON DELETE CASCADE,
    INDEX idx_status_jobs (status)
);

INSERT INTO employment_types (type) 
VALUES ('Contract'),('Internship'),('Full-Time'),('Temporary'),('Part-Time');

INSERT INTO position_level (position_levels) 
VALUES ('CEO/GM/Direktur/Manajer Senior'),('Manajer/Asisten Manajer'),('Supervisor/Koordinator'),('Pegawai non-manajemen & non-supervisor'),('Lulusan baru/Pengalaman kerja kurang dari 1 tahun');


INSERT INTO jobs (
	id, 
	title, 
	description,
	position_level_id,
	employment_type_id
) 
VALUES
	(UUID(), 'Database Administrator', '<p><strong><big>Job Description</big></strong><br/>
- Menguasai ms sql, sql cluster replication/ha, tuning query, function/store procedure<br/>
- Pengalaman: Minimun 2 tahun sebagai DBA</p>

<p><strong><big>Minimum Qualification</big></strong><br/>
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br/>
2 year(s)</p>

<p><strong><big>Skills</big></strong></p>

<ul>
	<li>SQL SERVER</li>
	<li>MY SQL</li>
</ul>

<p><strong><big>Benefit</big></strong><br />
<big>-</big></p>',1,1),
	(UUID(), 'Product Owner', '<p>&nbsp;</p>

<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>services firm</li>
	<li>Strong Consulting proven track record</li>
	<li>Strong Systems Integration proven track record</li>
	<li>Strong Complex Solutioning proven track record</li>
	<li>Strong Delivery proven track record</li>
</ul>

<p><strong><big>Minimum Experience</big></strong><br />
At least 5 years&rsquo; experience in management consulting and systems integration in a top tier professional</p>

<p><strong><big>Benefit</big></strong><br />
competitive salary, health insurance</p>
',2,2),
	(UUID(), 'Fullstack Developer', '<p><strong><big>Job Description</big></strong><br />
-&nbsp;Strong knowledge of front-end technologies, including HTML, CSS, JavaScript, and front-end frameworks (e.g., React, Angular, Vue.js).<br />
-Familiarity with back-end development using one or more programming languages like Python, Ruby, Java, or Node.js.</p>

<p><strong><big>Minimum Qualification</big></strong><br />
-&nbsp;Experience working with databases (e.g., MySQL, MongoDB, PostgreSQL) and proficiency in writing efficient queries.<br />
-&nbsp;Understanding of version control systems (e.g., Git) and collaborative development workflows.</p>

<p><strong><big>Minimum Experience</big></strong><br />
1 year</p>

<p><strong><big>Benefit</big></strong><br />
competitive sallary</p>
',5,4),
	(UUID(), 'PHP Developer', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>Prolific experience working with PHP, SQL, and JavaScript</li>
	<li>Demonstrable experience working with SOAP/REST</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
3 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',4,3),
	(UUID(), 'React Native Developer', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>Prolific experience working with React Native, SQL, and JavaScript</li>
	<li>Demonstrable experience working with SOAP/REST</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
3 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',5,3),
	(UUID(), 'Golang Developer', '<p>&nbsp;</p>

<p><big><strong>Position Summary</strong></big></p>

<p>As a Software Engineer (Golang), you design, implement, and maintain software solutions to meet given requirements. The ideal candidate should be able to guide others and provide exemplary solutions to complex problems.</p>

<p><br />
<strong><big>Job Description</big></strong></p>

<ul>
	<li>Review and provide feedback on your coworker&rsquo;s changes</li>
	<li>Investigate issues in any code to root cause and create fixes</li>
	<li>Design and develop software solutions to meet requirements</li>
	<li>Write code comments, change descriptions and documentation so others can maintain your work</li>
	<li>Create end-to-end tests that validate your code in context</li>
	<li>Develop and implement short-term and suggest long-term mitigations for incidents</li>
	<li>Train and mentor less experienced engineers</li>
	<li>Collaborate with other engineers and stakeholders to assist them when they face obstacles</li>
	<li>Maintain a high standard of engineering excellence by adhering to best practices</li>
	<li>Perform other duties as assigned</li>
</ul>

<p><strong><big>Requirements</big></strong></p>

<ul>
	<li>Specializes in distributed systems, experience with technologies like AWS, GoLang, Java, and Databases</li>
	<li>At least 4 year of professional software engineering experience</li>
	<li>Experience with the full lifecycle of a software product</li>
	<li>Professional experience working with multiple languages and technologies</li>
	<li>Real world experience making tradeoffs with algorithms, data structures and network protocols</li>
	<li>Proficiency in written and verbal English language to succeed in a remote work environment</li>
</ul>

<p><strong><big>Benefit</big></strong><br />
competitive salary, health insurance</p>
',4,3),
	(UUID(), 'IOS Developer', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>Proficient in Google Cloud, Jenkins, Docker and Terraform</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
	<li>Available to work full time and on-site</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
2 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',5,1),
	(UUID(), 'Project Support', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>Project Support</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
2 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',1,2),
	(UUID(), 'UI UX Designer', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>UI UX figma</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
2 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',2,3),
	(UUID(), 'DevOps', '<p><strong><big>Job Description</big></strong></p>

<ul>
	<li>Proficient in Google Cloud, Jenkins, Docker and Terraform</li>
	<li>Self-motivated, passionate &amp; eager to improve by keeping up-to-date with the latest trends in back-end development</li>
	<li>Good communication &amp; teamwork skills. Capable of collaborating with other team members</li>
	<li>Good self-managing skills, particularly with keeping up with timeline/deadline</li>
	<li>Available to work full time and on-site</li>
</ul>

<p>&nbsp;</p>

<p><strong><big>Minimum Qualification</big></strong><br />
Pendidikan terakhir D3/S1 Jurusan terkait</p>

<p><strong><big>Minimum Experience</big></strong><br />
2 year(s)</p>

<p><strong><big>Benefit</big></strong><br />
We offer a competitive salary, flexible working hours, and a fully remote work environment.&nbsp;</p>
',3,4);

select * from employment_types;
select * from position_levels;
select * from jobs;