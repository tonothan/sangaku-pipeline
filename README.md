# sangaku-pipeline

Throughout the Edo period, an idiosyncratic geometric discourse flourished in Japan, surviving in form of the so-called sangaku: votive tablets presenting geometric riddles in word and image, dedicated by private individuals in Buddhist temples and Shinto shrines. Around 1000 sangaku survive to date; spread out as they are throughout Japan's temples, however, their systematic exploration and a better understanding of their cultural function remains a challenge and the tradition as a whole practically unknown outside of Japan.
Our 3-year research project based at Waseda University, Tokyo, is an attempt to put sangaku on the map of the global history of mathematics: the first and primary objective is to digitize all sangaku surviving in wood and print and showcase them in a bilingual IIIF-based archive. 
There is an inherent challenge to the project: unlike most digitization initiatives pursued by libraries and cultural institutions, the primary material itself is not available in-house but held in temples spread out across Japan. Collecting decentralized material in such quantity, however, will require a collaborative effort. The project will therefore necessitate a workflow that links a user-friendly interface for submitting images to a pipeline streamlining the conversion of the material into IIIF manifests.

In this project, we would like to present a pilot for such a pipeline: from data collection to database management and the automated creation of IIIF manifests. 
The requirements are as follows: the submission of data should be possible from anywhere and user-friendly; there should be a window for control between the submission of data and its conversion to a IIIF manifest; the single steps of the process should be automated wherever feasible in order to limit manual intervention to a minimum. 
The pilot presented uses a Google Form for data submission and MySQL for database management; the orchestrator is written in Golang. With a view to the future, the final deliverable will be a docker container allowing the use of alternative interfaces or databases. 

## Database auth

To connect to the database create a `.env` file containing the following data
```
MONGOURI=mongodb+srv://{USERNAME}:{PASSWORD}@{DATABASE-NAME}.mongodb.net/?retryWrites=true&w=majority
```