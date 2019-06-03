# team-lgtm

```bash
$ gcloud beta run deploy --image gcr.io/team-lgtm-dev/backend --update-env-vars=`awk '{ORS=","} {print}' backend/env`
```