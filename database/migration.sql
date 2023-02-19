create table `Student`(
	Student_Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    Student_Name char(20),
    Course char(20),
    Department char(20),
    Place char(20)
);

Insert into `Student`
	(`Student_Name`,`Course`,`Department`,`Place`) values
		('Prashant','B.Tech','IT','Ghaziabad'),
        ('Udit','B.Tech','CSE','Ghaziabad');
