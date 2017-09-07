from robot.libraries.BuiltIn import BuiltIn
from selenium.webdriver.common.keys import Keys
import time
s2l = BuiltIn().get_library_instance('Selenium2Library')
#driver = s2l._current_browser()

"""
Call methods from s2l code, such as from
https://github.com/robotframework/Selenium2Library/blob/master/src/Selenium2Library/keywords/waiting.py
you could call:
s2l.wait_until_element_is_visible(locator)
"""

def lib_get_webdriver_instance():
	return s2l._current_browser() # ??? THIS WORKS ???

def lib_search_for_test():
	elem = s2l._current_browser().find_element_by_name("q")
	elem.clear()
	elem.send_keys("test")
	elem.send_keys(Keys.RETURN)

def lib_close_browser():
	s2l._current_browser().close()

def lib_sleep(duration):
	time.sleep(float(duration))

def lib_return_url():
	return s2l._current_browser().current_url

def lib_source_should_not_contain(text):
	assert text not in s2l._current_browser().page_source

def lib_try_type(element, text):
  s2l._current_browser().type(element, text)