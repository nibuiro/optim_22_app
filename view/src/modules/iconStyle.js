export function iconStyle(size, image) {
    const defaultIcon = "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4NCjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+PHN2Zw0KICAgIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHZlcnNpb249IjEuMSIgd2lkdGg9IjI0IiBoZWlnaHQ9IjI0Ig0KICAgIHZpZXdCb3g9IjAgMCAyNCAyNCI+DQogICAgPHBhdGggc3Ryb2tlPSJkaW1ncmF5IiBmaWxsPSJkaW1ncmF5Ig0KICAgICAgICBkPSJNMTIsNEE0LDQgMCAwLDEgMTYsOEE0LDQgMCAwLDEgMTIsMTJBNCw0IDAgMCwxIDgsOEE0LDQgMCAwLDEgMTIsNE0xMiwxNEMxNi40MiwxNCAyMCwxNS43OSAyMCwxOFYyMEg0VjE4QzQsMTUuNzkgNy41OCwxNCAxMiwxNFoiIC8+DQo8L3N2Zz4=";
    const noIcon = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII=";
    let icon = "";
    if (image === noIcon || image === "") {
        icon = defaultIcon;
    } else {
        icon = image;
    }
    return {
        width: `${size}px`,
        height: `${size}px`,
        backgroundImage: `url("${icon}")`,
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "center",
        borderRadius: "100%"
    };
}