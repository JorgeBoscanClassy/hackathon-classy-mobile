#  ******************************************************************************
#             Author:  Emad Borjian
#       Date Created:  04-27-2022
#        Description:  This feature file is intended to test scenarios related to 
#                      how the Launch Screen is rendered.
#  ******************************************************************************

Feature: Display of the Launch Screen

Background:
  # DEV ENVIRONMENT
  # END

  # STAGING ENVIRONMENT 
  # END

  # PROD ENVIRONMENT
  # END

  # GENERAL
  * def appName = 'Classy Mobile'
  * def companyName = 'Classy'
  * def color = 'White'
  # END

@P3
@Regression
@TestCaseKey=TE-TC-10003
@Route=exp://34-hbi.classyomid.classy-mobile.exp.direct:80
Scenario: Verify the Launch Screen is rendered correctly

  # Navigate to app icon
  Given an admin navigates to 'Home Screen' page where app icon is located
  
  # Launch the app
  When select appName
  
  # Verify Launch Screen is displayed and rendered correctly
  Then verify 'Launch Screen' is displayed
  And verify 'Launch Screen' contains companyName on color background