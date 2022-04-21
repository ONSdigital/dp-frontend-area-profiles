import("mapbox-gl");
import { MapComponent as _MapComponent } from "dp-maps-js";

interface IMapComponent {
    init: () => void;
}

export class MapComponent implements IMapComponent {
    public init(): void {
        const options = {
            style: "mapbox://styles/mapbox/cjaudgl840gn32rnrepcb9b9g",
            center: [-1.2471735, 50.8625412] as any,
            zoom: 6,
            token: process.env.MAPBOX_ACCESS_TOKEN,
            mapID: "map",
        }
        const map = new _MapComponent(options);
        map.init();
    }
}