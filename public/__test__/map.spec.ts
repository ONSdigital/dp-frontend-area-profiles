import "mapbox-gl";
import { IMapComponentOptions, MapComponent } from "dp-maps-js";

jest.mock("dp-maps-js", () => ({
    MapComponent: jest.fn().mockImplementation(() => {
        return {
            init: () => jest.fn(),
        }
    })
}));

describe("map", () => {
    describe("#MapComponent()", () => {
        test("#init()", () => {
            const options: IMapComponentOptions = {
                style: "",
                center: [-7.9454024125535625, 48.95006696529006],
                token: "",
                geoDataURL: "<GEO_DATA_URL>",
            };
            const mapComponent = new MapComponent(options);
            const mapComponentSpy = jest.spyOn(mapComponent, "init");
            mapComponent.init();
            expect(mapComponentSpy).toBeCalled();
        })
    });
});
