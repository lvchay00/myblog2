<html>
<head>
<title></title>
</head>
<body>
<form action="/AddClass/Add" method="post">

         <div class="col-md-12 column">
        <div class="col-md-2 column">

	  <input type="text" name="name"  style="padding-right:30px; width:150pt;height:30pt" >
          </div>
    <input type="submit" value="添加">	
          </div>
  
</form>
     <div class="col-md-12 column">
	 <p>       </p> 
	 </div>
<form action="/AddClass/Delete" method="post">
   
         <div class="col-md-12 column">
        <div class="col-md-2 column">
            <select name="sid" id="sid" class="form-control" style="padding-right:30px; width:150pt;height:30pt"  >	
            {{range .Sections}}
                <option value="{{.Id}}">{{.Name}}   </option>
              {{end}}
            </select>	
        </div>
	   <input type="submit" value="删除">	

          </div>

</form>

    <div class="col-md-12 column">
	 <p>       </p> 
	 </div>
	
<div class="col-md-12 column"><a href="/">返回主页</a></div>   
</body>
</html>