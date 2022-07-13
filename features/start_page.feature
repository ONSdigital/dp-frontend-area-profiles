Feature: Groups
    Scenario: GET /areas and checking the response status 200
    When I navigate to "/areas"
    Then the beta phase banner should be visible
    And the improve this page banner should be visible
    And element "[data-test='questions']" should be visible
    # -- List
    And the page should have the following content
        """
            {
                "[data-test='questions'] > li:nth-child(1)": "how many people live there",
                "[data-test='questions'] > li:nth-child(2)": "how crowded it is",
                "[data-test='questions'] > li:nth-child(3)": "people's average age",
                "[data-test='questions'] > li:nth-child(4)": "how many people think their general health is good",
                "[data-test='questions'] > li:nth-child(5)": "how many households where English is not the main language",
                "[data-test='questions'] > li:nth-child(6)": "how many households where Welsh is the main language (Wales only)",
                "[data-test='questions'] > li:nth-child(7)": "how many households are owned with a mortgage, loan or shared ownership"
            }
        """
    # -- Top section 
    And the page heading should be "Find facts and figures about areas in England or Wales"
    And the first paragraph should have a link of "England"
    And the first paragraph should have a second link of "Wales"
    # -- Country section
    And the country section sub heading is "Other countries"
    And the country section first paragraph contains link with text "areas in Scotland on Scotland’s Census website"
    And the country section second paragraph contains link with text "areas in Northern Ireland, see the Northern Ireland Statistics and Research Agency website"

