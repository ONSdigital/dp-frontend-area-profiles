Feature: Groups

    Scenario: Get /areas/{id} and checking the response status 200
    When I navigate to "/areas/E92000001"
    Then the beta phase banner should be visible
    And the improve this page banner should be visible
    And element "[data-test='svg-icon1']" should be visible
    And element "[data-test='hr1']" should be visible
    And the page should have the following content
        """
            {
                "h1.ons-u-fs-xxxl": "England",
                "[data-test='overview']": "Overview",
                "[data-test='p1']": "Facts and figures about people living in England",
                "[data-test='p2']": "Find out:",
                "[data-test='questions'] > li:nth-child(1)": "how many people live there",
                "[data-test='questions'] > li:nth-child(2)": "how crowded it is",
                "[data-test='questions'] > li:nth-child(3)": "people's average age",
                "[data-test='questions'] > li:nth-child(4)": "how many people think their general health is good",
                "[data-test='questions'] > li:nth-child(5)": "how many households where English is not the main language",
                "[data-test='questions'] > li:nth-child(6)": "how many households are owned with a mortgage, loan or shared ownership",
                "[data-test='ViewFactsFiguresNomis']" : "View facts and figures on Nomis "
            }
        """
    # -- Breadcrumbs
    And the page should contain "breadcrumbs" with list element text "Home,Areas" at 3 depth
    # -- Area relations    
    And the relations sub heading should be "Areas Within England"
    And the relations sections should have 3 external links
