$(function() {
  var current_filename
  var list          = $("#list")
  var logs          = $("#logs")
  var listTemplate  = _.template($("#list-template").html())
  var logsTemplate  = _.template($("#logs-template").html())
  var stream        = new EventSource(document.location.href)

  stream.onmessage = function(e) {
    line = JSON.parse(e.data)
    line.name = line.filename.slice(0, -4)

    if(line.filename != current_filename) {
      list.append(listTemplate(line))
      logs.append(logsTemplate(line))
      current_filename = line.filename
    }

    $("#" + line.name).find("code").append(line.text + "\n")
  }

  stream.onerror = function() {
    stream.close()
  }
})
