package main

import (
	"html/template"
	"net/http"
)

// RecipeData is a struct to hold recipe form data
type RecipeData struct {
	ClassName        string
	Name             string
	FirstIngredient  string
	SecondIngredient string
	Result           string
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form values from the request
	classname := r.FormValue("classname")
	name := r.FormValue("name")
	firstIngredient := r.FormValue("first_ingredient")
	secondIngredient := r.FormValue("second_ingredient")
	result := r.FormValue("result")

	// Create a RecipeData instance with form data
	recipeData := RecipeData{
		ClassName:        classname,
		Name:             name,
		FirstIngredient:  firstIngredient,
		SecondIngredient: secondIngredient,
		Result:           result,
	}

	// Parse the template
	tmpl, err := template.New("recipeTemplate").Parse(recipeTemplate)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the recipe data
	err = tmpl.Execute(w, recipeData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

const recipeTemplate = `
class {{ .ClassName }}_Recipe extends RecipeBase
{
	override void Init()
	{
		m_Name = "{{ .Name }}";
		m_IsInstaRecipe = false;
		m_AnimationLength = 1.5;
		m_Specialty = -0.02;

		m_MinDamageIngredient[0] = -1;
		m_MaxDamageIngredient[0] = 3;

		m_MinQuantityIngredient[0] = 4;
		m_MaxQuantityIngredient[0] = -1;

		m_MinDamageIngredient[1] = -1;
		m_MaxDamageIngredient[1] = 3;

		m_MinQuantityIngredient[1] = -1;
		m_MaxQuantityIngredient[1] = -1;

		InsertIngredient(0,"{{ .FirstIngredient }}");
		m_IngredientAddHealth[0] = 0;
		m_IngredientSetHealth[0] = -1;
		m_IngredientAddQuantity[0] = -4;
		m_IngredientDestroy[0] = false;
		m_IngredientUseSoftSkills[0] = false;

		InsertIngredient(1,"{{ .SecondIngredient }}");

		m_IngredientAddHealth[1] = 0;
		m_IngredientSetHealth[1] = -1;
		m_IngredientAddQuantity[1] = -10;
		m_IngredientDestroy[1] = false;
		m_IngredientUseSoftSkills[1] = false;

		AddResult("{{ .Result }}");

		m_ResultSetFullQuantity[0] = false;
		m_ResultSetQuantity[0] = -1;
		m_ResultSetHealth[0] = -1;
		m_ResultInheritsHealth[0] = -2;
		m_ResultInheritsColor[0] = -1;
		m_ResultToInventory[0] = -2;
		m_ResultUseSoftSkills[0] = false;
		m_ResultReplacesIngredient[0] = -1;
	}

	override bool CanDo(ItemBase ingredients[], PlayerBase player)
    {
        return true;
    };

	override void Do(ItemBase ingredients[], PlayerBase player, array<ItemBase> results, float specialty_weight)
	{
		Debug.Log("Recipe Do method called","recipes");
	};
};
`

func main() {
	http.HandleFunc("/recipe", recipeHandler)

	// Serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
