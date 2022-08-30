import { SVGRenderer } from "highcharts";
import "../sass/index.scss";
var Highcharts = require('highcharts'); 

const blue = "#206095";

enum Color {
    BLUE = "#206095",
    AQUA = "#00A3A6",
    GREEN = "#A8BD3A",
    GREY = "#E2E2E3",
}

const chartHeight = 500;
const barWidth = 24;
const groupPadding = 0.1; // %

const halfGroupPadding = chartHeight / groupPadding;

class Chart {
   public init() {
        document.addEventListener("DOMContentLoaded", () => {
            const chart = (Highcharts).chart("hcContainer", {
                chart: {
                    type: "bar",
                    marginLeft: 50, 
                },
                title: {
                    text: ""
                },
                yAxis: {
                    opposite: true,
                    min: 0,
                    minTickInterval: 25,
                    max: 100,
                    gridLineColor: '#222222',
                    gridZIndex: 4,
                    maxPadding:0,
                    endOnTick:false 
                },
                xAxis: {
                      labels: {
                        enabled: false,
                    },
                },
                plotOptions: {
                    bar: {
                        groupPadding: groupPadding,
                        borderWidth: 0,  
                        grouping: true,
                        pointWidth: barWidth,
                        // pointPadding: 0,
                        // pointPlacement: 0,
                        dataLabels: {
                            defer: false,
                            enabled: true,
                            overflow: 'allow',
                            align: "bottom",
                            verticalAlign: "middle",
                            x: -50,
                            y: 5,
                            color: "black",
                            crop: false,
                            className: "highcharts-poc__data-labels",
                            inside: true,
                            format: "{y} %",
                        },
                    },
                    // series: {
                    //     pointPadding: 0,
                    // }
                },
                series: [
                    {
                        data: [
                            { 
                                id: "bgBar",
                                y: 10,
                                color: Color.BLUE, // 1 Manchester
                            }, 
                            { 
                                y: 20,
                                color: Color.BLUE, // 4 Manchester
                            }, 
                            {
                                y: 30,
                                color: Color.BLUE, // 8 Manchester
                            }
                        ] 
                    },
                    {
                        data: [
                            { 
                                y: 40,
                                color: Color.AQUA, // 2 Stockport
                        
                            }, 
                            { 
                                y: 50,
                                color: Color.AQUA, // 2 Stockport
                            }, 
                            {
                                y: 60,
                                color: Color.AQUA, // 2 Stockport
                            },
                        ] 
                    },
                    {
                        data: [
                            { 
                                y: 70,
                                color: Color.GREEN, // 3 Rochdale
                        
                            }, 
                            { 
                                y: 80,
                                color: Color.GREEN, // 3 Rochdale
                            }, 
                            {
                                y: 90,
                                color: Color.GREEN, // 3 Rochdale
                            },
                        ] 
                    },
                ]
            }, (chart: any) => {
                // top 132
                    console.log("called : ", chart)
                    console.log("plot : ", chart.marginTop)
                    const series = chart.yAxis[0].series;
                    chart.renderer.rect(chart.plotLeft, chart.plotTop+chart.spacing[0], chart.plotWidth, barWidth, 0)
                        .attr({
                            fill: Color.GREY,
                            zIndex: -1
                        })
                        .add();
            }); 
        });
    }
}


const chart = new Chart();
chart.init();
