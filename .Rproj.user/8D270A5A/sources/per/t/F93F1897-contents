library(shiny)
library(shinydashboard)
library(googlesheets)
library(plotly)
library(stringr)

ui <- dashboardPage(skin = "red",
    dashboardHeader(title = tags$img(src = 'russia.png', height = '65%')),
    dashboardSidebar(collapsed = FALSE,
        sidebarMenu(
            menuItem("Home", tabName = "home", icon = icon("home")),
            menuItem("Estatisticas", tabName = "estat", icon = icon("bar-chart")),
            menuItem("Quem somos", tabName = "quem", icon = icon("address-card"))
        )
    ),
    dashboardBody(
        tabItems(
            tabItem(tabName = "home",
                fluidRow(
                    uiOutput("atualiza_data")
                )
            ),
            tabItem(tabName = "estat",
                plotlyOutput("grafico_estado"),
                br(),
                br(),
                fluidRow(
                   
                        valueBoxOutput("box_fila")   
                    
                )
            ),
            tabItem(tabName = "quem",
                box(
                    title = "Projeto R.U.SSIA",
                    status = "warning",
                    collapsible = FALSE,
                    h4(
                        "Somos o Projeto R.U.SSIA, criado por amigos durante a disciplina
                        de Projetos em Engenharia de Computacao na Universidade Federal de Sao Paulo
                        (UNIFESP)."
                    )
                ),
                box(
                    title = "Nossa missao",
                    status = "warning",
                    collapsible = FALSE,
                    h4(
                        "Com o objetivo de facilitar a vida dos estudantes de nossa universidade, que chegam
                        a passar horas nas filas do refeitorio universitario (RU), criamos esse aplicativo."
                    )
                ),
                box(
                    title = "Menos filas",
                    status = "warning",
                    collapsible = FALSE,
                    h4(
                        "Hoje buscamos diminuir o tempo gasto pelos alunos que utilizam da refeicao universitaria
                        devido a necessidade de ficar tempo integral na faculdade esse projeto foi desenvolvido, facilitando
                        a transicao de informacoes entre a NUTRIMENTA e os alunos por meio 
                        da disposicao dos dados sobre o quao cheio o RU se encontra."
                    )
                ),
                box(
                    title = "Equipe",
                    status = "warning",
                    collapsible = FALSE,
                    h5("Bruno Ogata"),
                    h5("Carolina Colla"),
                    h5("Jaime Ossada"),
                    h5("Joao Victor de Mesquita"),
                    h5("Pedro Naresi"),
                    h5("Raphael Faria")
                )
            )
        )
    )
)

server <- function(input, output) {
    tabela <- loadData()
    names(tabela) <- c("date", "hora", "estado", "nil") 
    last_data <- nrow(tabela)
    
    output$atualiza_data <- renderUI({
        print(tabela$date[last_data])
        box(title = paste("Fila do dia:", tabela$date[last_data]), status = "warning", collapsible = TRUE, 
            output$estado_fila <- renderValueBox({
                if(tabela$estado[last_data] == "cheio"){
                    valueBox("Cheio", "O R.U. parece estar um pouco lotado no momento, que tal ir daqui uns minutinhos?", icon = icon("thermometer-4"), color = "red")   
                } else {
                    valueBox("Vazio", "Corre pro R.U.!", icon = icon("thumbs-up"), color = "green")
                }
            }),paste("Atualizado ultima vez as:", tabela$hora[last_data])
            
        )
    })
    
    output$box_fila <- renderValueBox({
        tab <- table(tabela$estado[(nrow(tabela)-10):nrow(tabela)])
        count_cheio <- tab[names(tab) == "cheio"]
        media <- count_cheio * 30
        
        valueBox("Media de espera", paste(media, "segundos de fila cheia nos ultimos 5 minutos"), icon = icon("hourglass-half"), color = "green")
    })
   
    output$grafico_estado <- renderPlotly({
        estados <- tabela$estado
        estado_v1 <- str_replace_all(estados, "cheio", "1")
        estado_v1 <- str_replace_all(estado_v1, "vazio", "0")
        estados_v2 <- as.numeric(estado_v1)
        aux <- data.frame(hora = as.character(tabela$hora), estados = estados_v2)
        plot_ly(aux, x = ~hora, y = ~estados, type = 'scatter', mode = 'lines')
    })
}

loadData <- function(){
    sheet_key <- "11RcJjyf5nNYkO7EbhurNYZl8JccohGouigOJkjNh5UA"
    ss <- gs_key(sheet_key)
    tabela <- gs_read(ss = ss)
    return(as.data.frame(tabela))
}

# Run the application 
shinyApp(ui = ui, server = server)

