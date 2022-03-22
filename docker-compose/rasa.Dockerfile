FROM rasa/rasa:3.0.4-full
USER root
RUN pip install sanic==21.6.0
USER 1001