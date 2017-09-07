FROM ubuntu:16.04

WORKDIR workdir/
ADD . tests/

RUN apt-get update && apt-get --yes upgrade

RUN apt-get --yes install wget

RUN apt-get --yes install libxss1 libappindicator1 libindicator7 gconf-service libasound2 libgconf-2-4 libgtk-3-0 libnspr4 libxtst6 fonts-liberation libnss3 lsb-release xdg-utils
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN dpkg -i google-chrome*.deb

RUN apt-get --yes install python2.7 python-pip xvfb
RUN pip install --upgrade pip
RUN pip install robotframework-selenium2library robotframework-xvfb

RUN apt-get --yes install unzip
RUN wget https://chromedriver.storage.googleapis.com/2.31/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN chmod +x chromedriver
RUN mv chromedriver /usr/local/bin/

CMD ["robot", "tests/check_page_title.robot"]