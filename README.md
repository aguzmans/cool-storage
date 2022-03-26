# Cool Storage Project plan
We are a small company that offers could storage as a service. We have a more or less working product, but we want to also offer a new version of that product that costs much less by taking advantage of AWS Glacier deep storage. The idea is for people to be able to ether backup DIRECTLY to Glacier or put what they have in the regular Cloud storage we offer into Glacier directly from out interface. Your job will be the CORE API that will actually do the work of storing and retrieving from Glacier to local computer or to our cloud service.

The idea of cool storage plan to bring AWS S3 Glacier to end user’s hands and reduce the cost of doing business with us considerably. Think of cool storage as a proxy between NiHao Cloud who will actually save things to AWS Glacier.

# The problem
NiHao Cloud needs to get cheaper specially cheaper for backups and information that is not needed right away.

## Target customer of this systen:
1.	Existing customers of the cloud collaboration storage system who want to save in space costs.
1.	New Customers who need a cheap backup option for BIG bulks of data.
2.	Customers who need to keep information safe from Governments and third parties. (

## Security considerations:
1.	The files stored need to be compressed and encrypted (by US or by customer).
2.	If we create and APP that would directly store things in glacier; Can we secure AWS credentials on an app on the client side? 
3.	Can we create a secure api that does all the processing and proxies data to glacier; from NiHao Cloud Seafile OR directly from APP?

## Cost and Pricing considerations:
This would be viable if we can change ~2x the cost of glacier, but make it FULLY compatible with nihao Cloud. This would also be our own software and not directly related to seafile, but it need to talk to seafile for many activities. So going back to billing; we would change 2x storage + 2x retrival cost. The retrival cost could be bulk (free), Standard(medium cost) or Expedite (expensive).

## The solution
From Nihao Cloud we need to be able to copy or move Libraries, files and folders. The flow would be for copy, just make a backup to glacier. For move would be copy to glacier and then remove from NiHao Cloud.

# The options to restore should be
1.	Download now(to users local system; phone, laptop, pc, etc) or restore to NiHao Cloud. And then you should be able to choose or configure retrival; if bulk, standard or express.

# To access the user will
1.	Auth will Use LDAP
2.	We will use the ORGs from seafile and their user system.
3.	The org admin should be able to use these features or give other users access to do this.
4.	Still need to make our mind about if we will bill for the interactions with NiHao Cloud or not.

# Other general characteristics of the architecture
1.	Auth with LDAP.
2.	DB: should be CHEAP, Distributed, minumally eventually consistent. Should have information to make smart storage decisions. We should explore NoSQL like DDB, MongoDB and ES.
3.	Programming language will probably be golang or rust. I prefer Golang. Will use Boto and other AWS SDK for golang. The software needs to have unit tests and other tests. The final objective of the first iteration will be an API.The Db will store information enough to know where to store for a use based on geolocation. 
4.	We should think of programming the system to have plugable storages in the future if we found a cheaper form of storage comapred to AWS glacier. So the usage of boto should be in the form of a plugin that is activated. Then this smart system should know when to use one storage or other based on user request and auntomate everything that is not defined by the user.

There are 3 types of glacier, this software should aim at the cheaper options. Probably Deep archieve Glacier instead of Instant retrival and Flexible stprage graciar. This at least for version 1.0.

# API calls
a)	Auth (token/creds) -> Success/failure
b)	GetDirContentlist (user, searchCriteria) -> list of files.  This function will not talk to glacier directly but only to the DB.
c)	PutFileFolderInStorage(path, destination, user) -> Success/failure. This function will upload a file to glacier and store all the meta data needed to retrieve it; region, size, encryption method, Compression method, etc. 
a.	There will be need to have multipart upload here and make sure uploads are consistent.
d)	RetriveObject(Target,Priority,DestinationPath,user) -> An object successfully download to local(phone/laptop,etc) or to NiHaoCloud.
e)	ReportUsageMetrics(user, amount in Kb?) -> success/failure. It should probably report to our billing dashboard (dash.sesamedisk.com). The report should be about storage (every hour) or about new upload/download requests and execution from glacier to NiHao Cloud or to local.
f)	ChangeStorageMode(for the next iteration)

# Other features that are not part of the core API
a)	Login.
b)	Access definition dashboard. Where organization (Org) admin can define policies.
c)	Configure and execute Cronjob schedule for batch jobs to archieve every x time from NiHao Cloud if the files have not been accessed and they are bigger than X size, etc..
d)	Web interface to copy and more from local and from NiHao Cloud. When we say copy and move understand upload/download and delete.
e)	Add Phone app features for android and IOS.

# Prototype UI:
Login:
User:
Password:	

Clear or login

# Access definition dashboard:
a)	Add remove org members if the org is multi users.
b)	Create cronjobs -> kubernetes cronjobs.
c)	Backup or move a file or folder NOW.
d)	Restore files or folders now.

## Configure and execute Cronjob 
Preferably to Create cronjobs as kubernetes cronjobs running on a separate process that reports results and ends once the scheduled job is finished.
This feature would be to automatically backup or move files from NiHao Cloud Storage to the Cool Storage. Not sure how the interface should be, but we need; to be able to automate by age of last access, let say scan everything that is older than 2 month without being accessed and move or backup to the Cool Storage. This could be simply backup this folder with a new copy in the Cool Storage. We might also want to be able to have restore automated cronjobs, but it would a bit less important. This step will probably include the “Backup or move a file or folder NOW.” step.
## Backup or move a file or folder NOW.
File: from local or from NiHao Cloud
Destination:
Move or just Copy: Option and explain the difference.

Clear or Execute
