{% import "github.com/Jiang-Gianni/dbt/db" %}
{% import "fmt" %}
{% func TestOutput(qrList []*db.QueryResults) %}{% for _, qr := range qrList %}### **[{%s= qr.Name %}]({%s= qr.FileName %}#L{%s= qr.Line %})**
* **Query**: {%s= qr.Query %}
* **Args**: {%s= fmt.Sprintf("%s", qr.Args) %}
* **Error**: {%s= qr.Error %}
* **Columns**: {%s= fmt.Sprintf("%s", qr.Columns) %}
* **Results**: {%s= fmt.Sprintf("%s", qr.Rows) %}

{% endfor %}
{% endfunc %}