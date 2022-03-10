import "mapbox-gl";
import { MapComponent } from "../ts/components/map";
import { MapComponent as _MapComponent } from "maps-js";

jest.mock("maps-js", () => ({
    MapComponent: jest.fn().mockImplementation(() => {
        return {
            init: () => jest.fn(),
        }
    })
}));

describe("map", () => {
    describe("#MapComponent()", () => {
        test("#init()", () => {
            const mapComponent = new MapComponent();
            const mapComponentSpy = jest.spyOn(mapComponent, "init");
            mapComponent.init();
            expect(mapComponentSpy).toBeCalled();
        })
    });
});
