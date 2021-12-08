# dp-frontend-renderer migration

The `dp-frontend-renderer` is being deprecated and the files migrated to the relevant controller/service.

The services are likely to be in a transient state, where a file has been migrated but the affected service has not been released to production. At this point, the files will still be in the `dp-frontend-renderer` but pending removal.  

Due to the toml/localisation files being site wide, the sections migrated are listed. The files cannot be removed until the whole migration is complete.

Use the tables as a guide.

## Templates

| Files (assets/templates) | New location | Deleted from renderer  
| ----- | ------------ | ---------------------
| ../banners/cookies.tmpl | [dp-renderer/assets/templates/banners](https://github.com/ONSdigital/dp-renderer/tree/main/assets/templates/partials/banners) | Yes  
| ../banners/covid.tmpl | [dp-renderer/assets/templates/banners](https://github.com/ONSdigital/dp-renderer/tree/main/assets/templates/partials/banners) | No  
| All json-Id templates (../partials/json-ld) | [dp-renderer/assets/templates/partials/json-ld](https://github.com/ONSdigital/dp-renderer/tree/main/assets/templates/partials/json-ld) | No  
|  ../partials/breadcrumb.tmpl | [dp-renderer/assets/templates/partials/breadcrumb.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/partials/breadcrumb.tmpl) | No  
| ../partials/feedback.tmpl | [dp-renderer/assets/templates/partials/feedback.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/partials/feedback.tmpl) | No  
| ../partials/footer.tmpl | [dp-renderer/assets/templates/partials/footer.tmpl](https://github.com/ONSdigital/dp-renderer/tree/main/assets/templates/partials/footer.tmpl) | No  
| ../partials/gtm-data-layer.tmpl | [dp-renderer/assets/templates/partials/gtm-data-layer.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/partials/gtm-data-layer.tmpl) | No  
| ../partials/header.tmpl | [dp-renderer/assets/templates/partials/header.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/partials/header.tmpl) | No  
| ../partials/spinner.tmpl | [dp-renderer/assets/templates/partials/spinner.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/partials/spinner.tmpl) | No  
| ../error.tmpl | [dp-renderer/assets/templates/error.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/error.tmpl) | No  
| ../main.tmpl | [dp-renderer/assets/templates/main.tmpl](https://github.com/ONSdigital/dp-renderer/blob/main/assets/templates/main.tmpl) | No  
| All dataset landing page templates (../datasetLandingPage)  | [dp-frontend-dataset-controller/assets/templates](https://github.com/ONSdigital/dp-frontend-dataset-controller/tree/develop/assets/templates) | Yes  
| ../partials/release-alert.tmpl | [dp-frontend-dataset-controller/assets/templates/partials](https://github.com/ONSdigital/dp-frontend-dataset-controller/tree/develop/assets/templates/partials) | No
| ../cookies-preferences.tmpl | [dp-frontend-cookie-controller/assets/templates](https://github.com/ONSdigital/dp-frontend-cookie-controller/tree/develop/assets/templates) | Yes
| ../search | [dp-frontend-search-controller/assets/templates](https://github.com/ONSdigital/dp-frontend-search-controller/tree/develop/assets/templates) | Yes

## Toml

| Section | New location  
| ------- | ------------  
| Header | [dp-renderer/assets/locales](https://github.com/ONSdigital/dp-renderer/tree/main/assets/locales) |  
| Cookies | [dp-frontend-cookie-controller/assets/locales](https://github.com/ONSdigital/dp-frontend-cookie-controller/tree/develop/assets/locales) |
| Dataset landing page | [dp-frontend-dataset-controller/assets/locales](https://github.com/ONSdigital/dp-frontend-dataset-controller/tree/develop/assets/locales)  
