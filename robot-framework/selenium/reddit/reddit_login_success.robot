*** Settings ***
Library  Selenium2Library
Test Setup  Open browser and login to reddit
Test Teardown  Close browser

*** Variables ***
${PAGE_URL}  http://www.reddit.com/
${USERNAME}  1234hello3214
${PASSWORD}  123123

*** Test Cases ***
Page Should Contain Link to ${USERNAME} Post Page
    Wait Until Page Contains  ${USERNAME}
    Page Should Contain Link   ${USERNAME}


*** Keywords ***
Open Browser And Login To Reddit
    Open Browser    ${PAGE_URL}      browser=chrome
    Input Text      user        ${USERNAME}
    Input Text      passwd      ${PASSWORD}
    Click Button    login


    #Page Should Contain Element  .user > a:nth-child(1)
    #Wait for Condition   return window.jQuery.active == 0;
    #Page Should Contain Link   https://www.reddit.com/user/${USERNAME}/
    #Page Should Contain Link   link=${USERNAME}
    #Page Should Contain  ${USERNAME}
