import * as monaco from 'monaco-editor'
let nombreGeneral


const myEditor = monaco.editor.create(document.getElementById("editorArea"), {
	language: "javascript",
	automaticLayout: true,
  theme: "vs",
});

const btn = document.getElementById("lecDocV")
btn.onclick = ()=>{
  leerDoc()
}

function leerDoc(){
  const archivo = document.getElementById("formFile").files[0]
  if (!archivo){
    alert("Error archivo no seleccionado")
    return;
  }
  nombreGeneral = archivo.name
  const lector = new FileReader()
  lector.onload = () => {
    const contenido = lector.result
    myEditor.setValue(contenido)
  }
  lector.readAsText(archivo)
}

const btn2 = document.getElementById("archivoNombre")
btn2.onclick = ()=>{
  const newName = document.getElementById("nameDoc").value;
  if (!newName) {
    alert("Por favor ingresar un nombre")
    return
  }
  guardarArchivoDeTexto(myEditor.getValue(),newName)
}

const btn3 = document.getElementById("guardarNormal")
btn3.onclick = ()=>{
  if (!nombreGeneral) {
    guardarArchivoDeTexto(myEditor.getValue(),"text.swift")
  } else {
    guardarArchivoDeTexto(myEditor.getValue(),nombreGeneral)
  }
}

const guardarArchivoDeTexto = (contenido, nombre) => {
  const a = document.createElement("a");
  const archivo = new Blob([contenido], { type: 'text/plain' });
  const url = URL.createObjectURL(archivo);
  a.href = url;
  a.download = nombre;
  a.click();
  URL.revokeObjectURL(url);
}


const btn4 = document.getElementById("ejecutar")
btn4.onclick = ()=>{
  if (!myEditor.getValue()) {
    alert("Editor vacio")
    return
  }
  enviarStringAlEndpoint("http://127.0.0.1:8000/",myEditor.getValue())
}

async function enviarStringAlEndpoint(url, texto) {
  const envio = await fetch(url,{
    method:"POST",
    headers: {
      "Content-Type": "text/plain"
    },
    body : JSON.stringify({"content":texto}),
  })
  const response = await envio.json()
  let consola = response.Consola
  let errors = response.Errores
  let simbolos = response.Simbolos
  if(errors.length == 0){
    let tmpConsola = ""
    consola.forEach(element => {
      tmpConsola += element + "\n"
    });
    document.getElementById("textConsol").innerHTML=""
    document.getElementById("tablaError").innerHTML=""
    document.getElementById("textConsol").innerHTML=tmpConsola
    let contentSimbols = ``
    for(let elements in simbolos){
      let tmp = simbolos[elements]
      let tipTemp = retornarTipo(0)
      let tipVector = retornarTipo(0)
      let contN
      let contC = []
      console.log(tmp)
      if (tmp.Tipo === 9 ) {
        tipVector = retornarTipo(tmp.TipoVector)
        //obtener de resultados para vector
        for (let index = 0; index < tmp.Resultados.length; index++) {
          let element = tmp.Resultados[index];
          let tmp2 = obtenerResultados(tmp.TipoVector,element)
          console.log(tipVector,tmp2)
          contC.push(tmp2)
        }
      } else {
        tipTemp = retornarTipo(tmp.Tipo)
        tipVector = retornarTipo(0)
        contN = obtenerResultados(tmp.Tipo,tmp.Resultado)
      }
      contentSimbols += `
        <tr>
          <td>${tmp.Nombre}</th>
          <td>${tmp.Constante}</th>
          <td>${tipTemp}</th>
          <td>${tipVector}</th>
          <td>${contN}</th>
          <td>${contC}</th>
          <td>${tmp.Linea}</th>
          <td>${tmp.Columna}</th>
        </tr>
      `
    }
    document.getElementById("tablaSimbols").innerHTML=contentSimbols
  }else{
    const mapErrors = new Map()
    for (let index = 0; index < errors.length; index++) {
      const element = errors[index];
      let id = ""+element.Linea + element.Columna
      if (!mapErrors.has(id)) {
        mapErrors.set(id,element)
      }
    }
    let contError = ``
    mapErrors.forEach((valor,clave) => {
      contError += `
        <tr>
          <td>${valor.Details}</th>
          <td>${valor.Linea}</th>
          <td>${valor.Columna}</th>
        </tr>
      `
    })
    console.log(contError)
    document.getElementById("tablaError").innerHTML = contError
  }
}

function retornarTipo(numero){
  if (numero === 0) {
    return "NIL"  
  } else if (numero === 1) {
    return "Character"
  } else if (numero === 2) {
    return "Bool"
  } else if (numero === 3) {
    return "Integer"
  } else if (numero === 4) {
    return "Float"
  } else if (numero === 5) {
    return "String"
  } else if (numero === 6) {
    return "Break"
  } else if (numero === 7) {
    return "Continue"
  } else if (numero === 8) {
    return "Rango"
  } else if (numero === 9) {
    return "Vector"
  }else if (numero === 10) {
    return "Compuesto"
  }
  return "NIL"
}

function obtenerResultados(tipo,resultado){
  if (tipo === 0) {
    console.log("nil0")
    return "nil"
  } else if (tipo === 1) {
    console.log("nil1")
    return resultado.ValorC
  } else if (tipo === 2) {
    console.log("nil2")
    return resultado.ValorB
  } else if (tipo === 3) {
    console.log("nil3")
    return resultado.Valor
  } else if (tipo === 4) {
    console.log("nil4")
    return resultado.ValorF
  } else if (tipo === 5) {
    console.log("nil5")
    return resultado.ValorS
  } 
  console.log("nil6")
}
/*
          <td>${}</th>
          <td>${}</th>
          <td>${}</th>
          <td>${}</th>
*/