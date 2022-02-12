<!--
SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# RASA models

These are the raw models for the Aether Chatbot

> For the running system the RASA executor pulls the "trained" model by HTTP from a Model 
Server (a simple HTTP server). This Model Server can be created with `make rasa-model-server-docker` 
which trains a model from the raw files and saves it in to a TAR GZ file `20220101-000000-aether-trained.tar.gz`
which can be retrieved at runtime from `http://rasa-model-server:8080/models/20220101-000000-aether-trained`

When working locally, to spare having to install RASA on Python on your machine, you can run
the Docker version of RASA and load the model as a mounted folder

First train the model
```bash
docker run -u 1000 -v `pwd`/rasa-models/rasa/:/app rasa/rasa:3.0.4 train
```

Then run the trained model as a web server as:
```bash
docker run -u 1000 -p 127.0.0.1:5005:5005 -v `pwd`/rasa-models/rasa/:/app rasa/rasa:3.0.4 run
```

> Note: This assumes your uid=1000 on your system - the default is 1001 and you
> will get permission errors if this does not match you actual user ID

You can then interact with RASA through HTTP by following the instructions here
https://rasa.com/docs/rasa/connectors/your-own-website 