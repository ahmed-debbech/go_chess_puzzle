function setDraggingMode(drag){
    localStorage.setItem("drag_mode", drag)
}

function getDraggingMode(){
    if(localStorage.getItem("drag_mode") == null){
        return "1"
    }
    return localStorage.getItem("drag_mode")
}
