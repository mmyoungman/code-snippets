# Taken from https://groups.google.com/d/msg/robotframework-users/3FPSc9iEIdY/DkRlh8UZBQAJ

*** Settings ***
Library          Selenium2Library
Test Setup       Open Chrome for Headless Testing
Test Teardown    Close Browser

*** Test Cases ***
Test Chrome Headless Run
    Go To    http://google.com
    Capture Page Screenshot
    Input Text  id=lst-ib    It's alive!
    Capture Page Screenshot
    Submit Form  id=tsf
    Capture Page Screenshot

*** Keywords ***
Open Chrome for Headless Testing
    ${c_opts} =     Evaluate    sys.modules['selenium.webdriver'].ChromeOptions()    sys, selenium.webdriver
    Call Method    ${c_opts}   add_argument    headless
    Call Method    ${c_opts}   add_argument    disable-gpu
    Call Method    ${c_opts}   add_argument    no-sandbox
    Call Method    ${c_opts}   add_argument    window-size\=1024,768
    Create Webdriver    Chrome    crm_alias    chrome_options=${c_opts}