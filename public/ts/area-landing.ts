// Sass
import "../sass/index.scss";
// Ts imports
import("mapbox-gl");
import { IMapComponentOptions, IMapPadding, MapComponent} from "dp-maps-js";
import "../json/geoData1.json"; // TODO test

const padding: IMapPadding = {
    top: 50, bottom:25, left: 15, right: 5,
};
const options: IMapComponentOptions = {
    style: "mapbox://styles/mapbox/cjaudgl840gn32rnrepcb9b9g",
    center: [-1.2471735, 50.8625412] as any,
    zoom: 6,
    token: process.env.MAPBOX_ACCESS_TOKEN,
    mapID: "map",
    padding,
    debug: true,
    bounds: MapComponent.setBounds([-7.197902920595226, 49.80281964843661], [2.1939240880979014, 55.490707526676545])
}

const map = new MapComponent(options);
map.init();
