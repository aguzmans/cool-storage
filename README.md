# Cool Storage Project plan
 get your personal storage to cost lest by taking advantage of AWS Glacier deep storage. The idea is for people to be able to ether backup DIRECTLY to Glacier or put what they have in the regular Cloud storage we offer into Glacier directly. This CORE API will actually do the work of storing and retrieving from Glacier to local computer or to another cloud service.

The idea of cool storage plan to bring AWS S3 Glacier to end user’s hands and reduce the cost of doing business considerably. Think of cool storage as a proxy between backup and regular local files and clouds.

# The problem
Get cheaper storage specially cheaper for backups and information that is not needed right away.

## Target audiance of this systen:
1.	People who need backups at the low cost of AWS glacier.

## Security considerations:
1.	The files stored need to be compressed and encrypted (by US or by user).
2.	Can we create a secure api that does all the processing and proxies data to glacier

# To access the user will
1.	Auth will Use LDAP
2.	We will use the ORGs from seafile and their user system.
3.	The org admin should be able to use these features or give other users access to do this.
4.	Still need to make our mind about if we will bill for the interactions with NiHao Cloud or not.

# Other general characteristics of the architecture
1.	Auth with LDAP.
2.	DB: will be MySQL.
3.	Programming language will be golang. Will use Boto and other AWS SDK for golang. 
4. The software needs to have unit tests and other tests. The final objective of the first iteration will be an API. 
5. The Db will store information enough to know where to store for a use based on geolocation. 
6.	We should think of programming the system to have plugable storages in the future if we found a cheaper form of storage comapred to AWS glacier. 
7. So the usage of glacier or other storages should be in the form of a plugin that is activated. 
8. Then this smart system should know when to use one storage or other based on user request and auntomate everything that is not defined by the user.
9. The system should be able to consider one storage class a cache for another. s3 as cache for Glacier and similar.

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
