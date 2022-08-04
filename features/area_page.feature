# Area Type field must map to the assets/locales/service.en.toml & assets/locales/service.cy.toml
# area type TOML table names in features/area_profiles_component.go areas.AreaDetails struct.

Feature: Groups
    Scenario: Get /areas/{id} and checking the response status 200
    When I navigate to "/areas/E92000001"
    Then the beta phase banner should be visible
    And the improve this page banner should be visible
    And element "[data-test='svg-icon1']" should be visible
    And element "[data-test='hr1']" should be visible
    And the area type should be "Country:"
    And the page should have the following content
        """
            {
                "h1.ons-u-fs-xxxl": "England Profile",
                "[data-test='overview']": "Overview",
                "[data-test='p1']": "Facts and figures from Census 2021 results on the population of England",
                "[data-test='p2']": "You can view and compare data on England with other areas about the topics:",
                "[data-test='questions'] > li:nth-child(1)": "population",
                "[data-test='questions'] > li:nth-child(2)": "identity",
                "[data-test='questions'] > li:nth-child(3)": "housing",
                "[data-test='questions'] > li:nth-child(4)": "work",
                "[data-test='questions'] > li:nth-child(5)": "education",
                "[data-test='questions'] > li:nth-child(6)": "health"
            }
        """
    And the Nomis link should point to "https://www.nomisweb.co.uk/reports/localarea?compare=E92000001"
    # -- Breadcrumbs
    And the page should contain "breadcrumbs" with list element text "Home,Areas" at 3 depth
    # -- Area relations    
    # And the relations sub heading should be "Areas Within England"
    And the relations sections should have 3 external links
    # -- Links
    And the first link text value should be "North East"
    And the second link text value should be "North West"
    And the third link text value should be "Yorkshire and The Humbe"
    And the first link href value should be "/areas/E12000001"
    And the second link href value should be "/areas/E12000002"
    And the third link href value should be "/areas/E12000003"
    # -- Related section
    And the Related content heading should be "Related content"
    And the page should have the following content
    """
        {
            "[data-test='related-ul'] > li:nth-child(1) > a": "Census"
        }       
    """