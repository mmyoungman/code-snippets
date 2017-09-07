*** Settings ***
Library  Selenium2Library
Library     XvfbRobot 
Test Setup  Open browser and go to page
Test Teardown  Close browser

*** Variables ***
${PAGE_URL}  file://${CURDIR}/page.html
${BROWSER}  chrome

*** Test Cases ***
Page Should Show Header
    [Documentation]  When visiting page, it should show "Hello World" text
    Page Should Contain  Hello World

*** Keywords ***
Open Browser And go to page
    Start Virtual Display    1920    1080
    ${chrome_options}=    Evaluate    sys.modules['selenium.webdriver'].ChromeOptions()    sys
    Call Method   ${chrome_options}   add_argument  --no-sandbox
    Create Webdriver    Chrome    my_alias    chrome_options=${chrome_options}   executable_path=/usr/local/bin/chromedriver
    Go to     ${PAGE_URL}
    #Open Browser   ${PAGE_URL}   browser=${BROWSER}   #desired_capabilities=--no-sandbox
    Set Window Size    1920    1080