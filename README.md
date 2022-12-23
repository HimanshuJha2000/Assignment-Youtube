# Assignment-Youtube

## Project Goal

To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

### Basic Requirements:

- [X] Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
- [X] A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
- [X] A basic search API to search the stored videos using their title and description.
- [X] Dockerize the project.
- [X] It should be scalable and optimised.

### Bonus Points:

- [X] Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
- [ ] Make a dashboard to view the stored videos with filters and sorting options (optional)

## Setup using Docker

```  
  # Clone the repository
  $ git clone git@github.com:HimanshuJha2000/Assignment-Youtube.git
  $ cd Assignment-Youtube/
  $ docker-compose -f docker-compose.dev.yml build
  $ docker-compose -f docker-compose.dev.yml up
  # This will start postgres, youtube-api and youtube-worker containers
```  
  
### Working

- There are two types of containers, i.e. youtube-api and youtube-worker. Youtube-api hold all the relevant APIs like adding new API key to the DB, Search API, Paginated response of videos API.
- While worker performs the task of running a cron asynchronously and fetches the youtube video data using Youtube API and saving it to the database.
- Relevant logs and error handling has been done.
- Support for auto switch over to new Youtube API key once the older key has been exhausted is added.
- Composite indexing has been done on title and description of a youtube video.

### API's Postman ScreenShots

- POST API to add a new Youtube Data API Key in the database
<img width="1011" alt="image" src="https://user-images.githubusercontent.com/35391094/209400299-0eeab7a0-50de-4d42-a9b7-ed911eb3f338.png">

- A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
<img width="1009" alt="image" src="https://user-images.githubusercontent.com/35391094/209400310-ac67ffe6-8605-4eeb-bd24-5a9338c7b355.png">

- Search API to search the stored videos using their title and description. (Shown by title)
<img width="1008" alt="image" src="https://user-images.githubusercontent.com/35391094/209400320-573d917e-92be-4aa8-81b7-30b8a9fe7f38.png">

  
