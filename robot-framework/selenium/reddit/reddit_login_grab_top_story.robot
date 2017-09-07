*** Settings ***
Library  Selenium2Library
Library  PrintVariable.py
Test Setup  Open browser and login to reddit
Test Teardown  Close browser

*** Variables ***
${PAGE_URL}  http://www.reddit.com/
${USERNAME}  1234hello3214
${PASSWORD}  123123
${TOPSTORY}  My grandma smoking
${SECONDSTORY}   The Veiled Virgin

*** Test Cases ***
#Page Should Contain Link to ${USERNAME} Post Page
#    Wait Until Page Contains  ${USERNAME}
#    Page Should Contain Link   ${USERNAME}

#Top Story Should Be ${TOPSTORY}
#    #Element Should Contain    css=div[id^=thing_t3]:nth-child(1)   ${TOPSTORY}
#    Element Should Contain    xpath=//div[contains(@id, 'thing_t3')][1]  ${TOPSTORY}

#Second Story Should Be ${SECONDSTORY}
#    #NOTE: nth-child(3) (and nth-of-type(3)) due to "<div class="clearleft"></div>"
#    #Element Should Contain    css=#siteTable div[id^=thing_t3]:nth-child(3)   ${SECONDSTORY}
#    Element Should Contain    xpath=//div[starts-with(@id, 'thing_t3')][2]   ${SECONDSTORY}

#Should Have 25 Stories
#    #49th div because 25 thing_t3 divs and 24 clearleft divs
#    #Page Should Contain Element    css=#siteTable div[id^=thing_t3]:nth-child(49)
#    Page Should Contain Element    xpath=//div[starts-with(@id, 'thing_t3')][25]

#Should Have 25 Stories Version 2
#	#Locator Should Match X Times   css=[id^=thing_t3]   25
#	Locator Should Match X Times   xpath=//div[starts-with(@id, 'thing_t3')]   25

#Top Story Should Have Over 10k Upvotes
#    #${UPVOTES}  Get Element Attribute  css=[id^=thing_t3]:nth-child(1) [class*="score likes"]@title
#    ${UPVOTES}  Get Element Attribute  xpath=//div[starts-with(@id, 'thing_t3')][1]//div[contains(@class, 'score likes')]@title
#    Should Be True   ${UPVOTES} > 10000

#Enter Preferences, Change NumSites to 25 via List
#	Wait Until Page Contains  ${USERNAME}
#	Click Link   preferences
#	Wait Until Page Contains  interface language
#	#Select From List   css=[name="numsites"]   25
#	Select From List   xpath=//select[@name="numsites"]   25
#	Click Button	save options
#	Wait Until Page Contains   interface language
#	#${CURRENTNUMSITES}   Get Value   css=[name="numsites"] [selected="selected"]
#	${CURRENTNUMSITES}   Get Value   xpath=//select[@name="numsites"]//option[@selected="selected"]
#	Should Be Equal   ${CURRENTNUMSITES}   25

#Enter Preferences, Disable Private RSS Feeds via Checkbox
#	Wait Until Page Contains  ${USERNAME}
#	Click Link   preferences
#	Wait Until Page Contains  interface language
#	Unselect Checkbox   private_feeds   #Before was using "css=[id="private_feeds"]" which also works
#	Unselect Checkbox   jquery=[id="private_feeds"]
#	Unselect Checkbox   xpath=//input[@id='private_feeds']
#	Click Button	save options
#	Wait Until Page Contains   interface language
#	Checkbox Should Not Be Selected   private_feeds

#Enter Preferences, Disable Thumbnails via Radio Button
#	Wait Until Page Contains  ${USERNAME}
#	Click Link   preferences
#	Wait Until Page Contains  interface language
#	# Radio Buttons locators are <input name="group"> and <input value="value">
#	Select Radio Button   media   off
#	Click Button	save options
#	Wait Until Page Contains   interface language
#	Radio Button Should Be Set To   media   off  

*** Keywords ***
Open Browser And Login To Reddit
    Open Browser    ${PAGE_URL}
    Input Text      user        ${USERNAME}
    Input Text      passwd      ${PASSWORD}
    Click Button    login
