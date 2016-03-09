<div id="content">
<h1>Registrazione Dati Veicolo</h1>
&nbsp;

{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}{{if .flash.notice}}
<h3>{{.flash.notice}}</h3>
&nbsp;
{{end}}
<form method="POST">
<table>


<tr>
    <td>Anno Immatricolazione: </td>
    <td><input name="annoimmatricolazione" type="text" value="{{.VeicoliDT.AnnoImmatricolazione}}" autofocus /> {{if .Errors.Targa}}{{.Errors.Targa}}{{end}}</td>
</tr>
<tr>
    <td>N° Libretto:</td>
    <td><input name="nlibretto" type="text" value="{{.VeicoliDT.NLibretto}}" /></td>
</tr>
<tr>
    <td>N° Telaio: </td>
    <td><input name="ntelaio" type="text" value="{{.VeicoliDT.NTelaio}}" /> {{if .Errors.DataFineFlotta}}{{.Errors.DataFineFlotta}}{{end}}</td>
</tr>
<tr>      
    <td>Marca: </td>
    <td><input name="marca" type="text" value="{{VeicoliDT.Marca}}" /> {{if .Errors.Note}}{{.Errors.Marca}}{{end}}</td>
</tr>
<tr>      
    <td>Modello: </td>
    <td><input name="modello" type="text" value="{{.VeicoliDT.Modello }}" /> {{if .Errors.Modello}}{{.Errors.Modello}}{{end}}</td>
</tr>
<tr>      
    <td>Normativa Euro: </td>
    <td><input name="noreuro" type="text" value="{{.VeicoliDT.NorEuro}}" /> {{if .Errors.NorEuro}}{{.Errors.NorEuro}}{{end}}</td>
</tr>
<tr>      
    <td>Kw: </td>
    <td><input name="kw" type="text" value="{{.VeicoliDT.Kw}}" /> {{if .Errors.Kw}}{{.Errors.Kw}}{{end}}</td>
</tr>
<tr>      
    <td>Cilindrata: </td>
    <td><input name="cilindrata" type="text" value="{{.VeicoliDT.Cilindrata}}" /> {{if .Errors.Cilindrata}}{{.Errors.Cilindrata}}{{end}}</td>
</tr>
<tr>      
    <td>ConsumoTeorico: </td>
    <td><input name="consumoteorico" type="text" value="{{.VeicoliDT.ConsumoTeorico}}" /> {{if .Errors.ConsumoTeorico}}{{.Errors.ConsumoTeorico}}{{end}}</td>
</tr>
<tr>      
    <td>KmAnno: </td>
    <td><input name="kmanno" type="text" value="{{.VeicoliDT.KmAnno}}" /> {{if .Errors.KmAnno}}{{.Errors.KmAnno}}{{end}}</td>
</tr>
<tr>      
    <td>CostoKm: </td>
    <td><input name="costokm" type="text" value="{{.VeicoliDT.CostoKm}}" /> {{if .Errors.CostoKm}}{{.Errors.CostoKm}}{{end}}</td>
</tr>
<tr>      
    <td>Pneumatici: </td>
    <td><input name="pneumatici" type="text" value="{{.VeicoliDT.Pneumatici}}" /> {{if .Errors.Pneumatici}}{{.Errors.Pneumatici}}{{end}}</td>
</tr>
<tr>      
    <td>Carburante: </td>
    <td><input name="carburante" type="text" value="{{.VeicoliDT.Carburante}}" /> {{if .Errors.Carburante}}{{.Errors.Carburante}}{{end}}</td>
</tr>
<tr><td>&nbsp;</td></tr>
<tr>
    <td>&nbsp;</td><td><input type="submit" value="RegisterVeicleDT" /></td>
</tr>
</table>
</form>
</div>




